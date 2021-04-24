package main

const sqlGetResults = `
SELECT
	padded_id id,
	name,
	CASE 	
		WHEN ts_headline(classification, query) LIKE '%<b>%' 					
				THEN CONCAT('Classification: ', ts_headline(classification, query))
		WHEN ts_headline(name, query) LIKE '%<b>%' 								
				THEN CONCAT('Name: ', ts_headline(name, query))
		WHEN ts_headline(padded_id, query) LIKE '%<b>%' 					
				THEN CONCAT('Pokedex: ', ts_headline(padded_id, query))
		WHEN ts_headline(array_to_string(types, ', '), query) LIKE '%<b>%'		
				THEN CONCAT('Types: ', ts_headline(array_to_string(types, ', '), query))
		WHEN ts_headline(array_to_string(abilities, ', '), query) LIKE '%<b>%' 	
				THEN CONCAT('Abilities: ', ts_headline(array_to_string(abilities, ', '), query))
	ELSE '' END AS snippet
FROM (
	SELECT
		padded_id, 
		name, 
		types, 
		classification, 
		abilities, 
		ts_rank(tsv, query) as rank, 
		query
	FROM
		pokemon, plainto_tsquery($1) query
	WHERE
		tsv @@ query
	ORDER BY
		rank DESC
	LIMIT 10
) AS results
ORDER BY
	rank DESC;
`

const sqlGetDetails = `
SELECT
	padded_id id,
	name,
	types,
	classification,
	abilities
FROM pokemon
WHERE padded_id=$1
`
