select es.name Employee
from Employee e
    inner join Employee es on es.managerId = e.id
where es.salary > e.salary;
