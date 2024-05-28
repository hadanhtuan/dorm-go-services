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