DELETE FROM Person
WHERE id NOT IN (
        SELECT min_id
        FROM (
                SELECT MIN(id) min_id
                FROM Person
                GROUP BY email
            ) tmp
    );
-- 以下のようにpersonテーブルを2つ別名つけるともっと早い
-- delete p1 from person p1,person p2
-- where p1.email=p2.email and p1.id>p2.id;
