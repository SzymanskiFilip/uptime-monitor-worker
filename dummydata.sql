DO $$
DECLARE 
    i INT := 0;
    j INT;
    current_date DATE := CURRENT_DATE;
    is_success BOOLEAN;
BEGIN
    WHILE i < 60 LOOP
        j := 0;
        WHILE j < 20 LOOP
            is_success := (RANDOM() > 0.5)::BOOLEAN;
            INSERT INTO statistics (url_id, headers, success, response_time, saved_at)
            VALUES (
                '9f814514-0ada-40e1-b7d8-cf0ff7f94467',
                'Header ' || (RANDOM() * 1000)::INT,
                is_success,
                CASE 
                    WHEN is_success THEN (RANDOM() * 4900 + 100)::INT
                    ELSE 0
                END,
                current_date - i
            );
            j := j + 1;
        END LOOP;
        i := i + 1;
    END LOOP;
END $$;
