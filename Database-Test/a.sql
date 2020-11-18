select c.customer_id, c.customer_name, count(o2.customer_id) as totalOrder from customers c, orders o2
where o2.customer_id = c.customer_id group by c.customer_id, o2.customer_id