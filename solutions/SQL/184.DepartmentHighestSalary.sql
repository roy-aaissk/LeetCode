WITH RankedSalaries AS (
        SELECT
            e.name AS Employee,
            e.salary AS Salary,
            d.name AS Department,
            DENSE_RANK() OVER (
                PARTITION BY e.departmentId
                ORDER BY e.salary DESC
            ) AS rn
        FROM Employee e
            JOIN Department d ON e.departmentId = d.id
        Order BY e.id DESC
    )

SELECT
    Department,
    Employee,
    Salary
FROM RankedSalaries
WHERE rn = 1;
