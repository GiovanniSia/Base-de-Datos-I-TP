-- Constraints tabla cliente
alter table cliente drop constraint cliente_pk;

-- Constraints tabla tarjeta
alter table tarjeta drop constraint tarjeta_pk;
alter table tarjeta drop constraint tarjeta_nrocliente_fk;

-- Constraints tabla comercio
alter table comercio drop constraint comercio_pk;

-- Constraints tabla compra
alter table compra drop constraint compra_pk;
alter table compra drop constraint compra_nrotarjeta_fk ;
alter table compra drop constraint compra_nrocomercio_fk ;

-- Constraints tabla rechazo
alter table rechazo drop constraint rechazo_pk ;

alter table rechazo drop constraint rechazo_nrocomercio_fk ;

-- Constraints tabla cierre  
alter table cierre drop constraint cierre_pk ;

-- Constraints tabla cabecera
alter table cabecera drop constraint cabecera_pk ;
alter table cabecera drop constraint cabecera_nrotarjeta_fk ;

-- Constraints tabla detalle 
alter table detalle drop constraint detalle_pk ;
alter table detalle drop constraint detalle_nroresumen_fk ;

-- Constraints tabla alerta
alter table alerta drop constraint alerta_pk ;

