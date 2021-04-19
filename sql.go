package main

const sqlGetResults = `
		SELECT
				id,
				name,
				CASE 	
						WHEN ts_headline(classification, q) LIKE '%<b>%' 					
								THEN CONCAT('Classification: ', ts_headline(classification, q))
						WHEN ts_headline(name, q) LIKE '%<b>%' 								
								THEN CONCAT('Name: ', ts_headline(name, q))
						WHEN ts_headline(CAST(id AS TEXT), q) LIKE '%<b>%' 					
								THEN CONCAT('Pokedex: ', ts_headline(CAST(id AS TEXT), q))
						WHEN ts_headline(array_to_string(types, ', '), q) LIKE '%<b>%'		
								THEN CONCAT('Types: ', ts_headline(array_to_string(types, ', '), q))
						WHEN ts_headline(array_to_string(abilities, ', '), q) LIKE '%<b>%' 	
								THEN CONCAT('Abilities: ', ts_headline(array_to_string(abilities, ', '), q))
				ELSE '' END AS snippet
		FROM (
				SELECT
					id, name, types, classification, abilities, ts_rank(tsv, q) as rank, q
				FROM
					pokemon, plainto_tsquery($1) q
				WHERE
					tsv @@ q
				ORDER BY
					rank DESC
			LIMIT 10
		) AS a
		ORDER BY
			rank DESC;
`

const sqlGetDetails = `
		SELECT
				id,
				name,
				types,
				classification,
				abilities
		FROM pokemon
		WHERE id=$1
`
