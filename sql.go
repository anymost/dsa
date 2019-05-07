package main

/**
select a.Name as Employee from Employee as a, Employee as b where a.ManagerId = b.Id && a.Salary > b.Salary;
select distinct a.Email from Person as a, Person as b where a.Email = b.Email && a.Id <> b.Id;

select Name
from (
    select a.Name as Name,
       b.Id as OrderId
      from Customers as a
          left join
          Orders as b
          on a.Id = b.CustomerId) as NOI
where OrderId is not null;


 */
