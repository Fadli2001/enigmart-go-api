package utils

const (
	INSERT_PRODUCT = `insert into m_product (id,name,price,stock)values(:id,:name,:price,:stock)`
	SELECT_PRODUCT_ID = `select id,name,price,stock from m_product where id = $1`
	SELECT_PRODUCT_LIST = `select id,name,price,stock,created_at,updated_at from m_product limit $1 offset $2`
	DELETE_PRODUCT = `delete from m_product where id = $1`
	UPDATE_PRODUCT = `update m_product set name=:name, price=:price, stock=:stock where id =:`

)