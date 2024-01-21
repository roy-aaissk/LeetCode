SELECT
    CASE
        WHEN COUNT(num) = 1 THEN MAX(num)
        ELSE NULL
    END AS num
FROM MyNumbers
GROUP BY num
ORDER BY num DESC
LIMIT 1;
