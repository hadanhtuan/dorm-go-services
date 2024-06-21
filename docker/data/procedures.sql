--Generate property_amenity
CREATE OR REPLACE PROCEDURE ADD_PROPERTY_AMENITY()
LANGUAGE plpgsql
AS $$ 
DECLARE
    prop RECORD;
    amen RECORD;
BEGIN
    -- Iterate over each property
    FOR prop IN SELECT * FROM property LOOP
        -- For each property, iterate over each amenity
        FOR amen IN SELECT * FROM amenity ORDER BY random() LIMIT 20 LOOP
            -- Insert a combination of property and amenity into property_amenity
            INSERT INTO property_amenity (property_id, amenity_id)
            VALUES (prop.id, amen.id);
        END LOOP;
    END LOOP;
END;
$$;

CALL ADD_PROPERTY_AMENITY();


--Generate reviews
CREATE OR REPLACE FUNCTION get_random_review_texts() RETURNS TEXT[] AS $$
BEGIN
    RETURN ARRAY[
        'Great place!', 
        'Had a wonderful time.',
        'Could be better.',
        'Not satisfied.',
        'Highly recommended!',
        'Will visit again.',
        'Average experience.',
        'Fantastic location.',
        'Service was excellent.',
        'Room was clean and spacious.',
        'Not worth the price.',
        'Will definitely return.',
        'Breakfast was delicious.',
        'Too noisy at night.',
        'Very comfortable stay.',
        'Staff were very helpful.',
        'Amenities were top-notch.',
        'Would not recommend.',
        'A home away from home.',
        'Loved the decor.',
        'Parking was a hassle.',
        'Great for families.',
        'Perfect for a weekend getaway.',
        'Air conditioning was poor.',
        'Check-in was smooth and easy.'
    ];
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE PROCEDURE generate_random_reviews()
LANGUAGE plpgsql
AS $$
DECLARE
    property_id TEXT;
    member_id TEXT;
    review_texts TEXT[];
    selected_texts TEXT[];
    rating INT;
    review_text TEXT;
    i INT;
BEGIN
    -- Get the array of review texts
    review_texts := get_random_review_texts();

    FOR property_id IN SELECT id FROM property LOOP
        -- Randomly select 7 unique review texts for each property
        selected_texts := ARRAY(
            SELECT review_texts[ceil(random() * array_length(review_texts, 1))]
            FROM generate_series(1, 7)
        );

        FOR i IN 1..7 LOOP
            rating := FLOOR(1 + RANDOM() * 5);  -- Generate a random rating between 1 and 5
            review_text := selected_texts[i];
            member_id := (SELECT id FROM member ORDER BY random() LIMIT 1); -- Select a random member
            INSERT INTO review (property_id, user_id, comment, overall_rating)
            VALUES (property_id, member_id, review_text, rating);
        END LOOP;
    END LOOP;
END;
$$;