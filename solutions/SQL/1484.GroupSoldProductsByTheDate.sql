SELECT
    sell_date,
    Count(DISTINCT product) num_sold,
    GROUP_CONCAT(
        DISTINCT product
        ORDER By product
    ) products
FROM Activities
GROUP BY sell_date
Order By sell_date asc;
