select p.product_name, p.basic_price, p.product_name, count(o.order_detail_id) as numberOfOrder from products p, orderdetails o where
        p.product_id = o.product_id group by p.product_id, o.product_id order by count(o.order_detail_id) desc