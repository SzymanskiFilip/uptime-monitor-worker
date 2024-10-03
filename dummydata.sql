/*the id on line 13 is the id of a URL*/

DO $$
DECLARE 
    i INT := 0;
    j INT;
    current_date DATE := CURRENT_DATE;
BEGIN
    WHILE i < 7 LOOP
        j := 0;
        WHILE j < 20 LOOP
            INSERT INTO statistics (url_id, headers, success, response_time, saved_at)
            VALUES (
                '8b91d2b2-daef-4e75-94ef-413e7111562e',
                'Header ' || (RANDOM() * 1000)::INT,
                (RANDOM() > 0.5)::BOOLEAN,
                (RANDOM() * 4900 + 100)::INT,
                current_date - i
            );
            j := j + 1;
        END LOOP;
        i := i + 1;
    END LOOP;
END $$;
