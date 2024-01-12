WITH ActivityWithPrevious AS (
    SELECT
        machine_id,
        process_id,
        activity_type,
        timestamp,
        LAG(timestamp) OVER (PARTITION BY machine_id, process_id ORDER BY timestamp) AS prev_timestamp
    FROM
        Activity
)
SELECT
    machine_id,
    ROUND(SUM(CASE WHEN activity_type = 'end' THEN timestamp - prev_timestamp ELSE 0 END) / COUNT(DISTINCT process_id), 3) AS processing_time
FROM
    ActivityWithPrevious
WHERE
    activity_type = 'end'
GROUP BY
    machine_id;

-- こんな方法がある
select distinct a1.machine_id, round(avg(a2.timestamp - a1.timestamp),3) as processing_time
from Activity a1 join Activity a2 on a1.machine_id = a2.machine_id and a1.process_id = a2.process_id and a1.activity_type = 'start' and a2.activity_type = 'end'
group by a1.machine_id
