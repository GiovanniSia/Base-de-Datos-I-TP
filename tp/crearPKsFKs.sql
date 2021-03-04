-- Constraints tabla cliente
alter table cliente add constraint cliente_pk primary key (nrocliente);

-- Constraints tabla tarjeta
alter table tarjeta add constraint tarjeta_pk primary key (nrotarjeta);
alter table tarjeta add constraint tarjeta_nrocliente_fk foreign key (nrocliente) references cliente (nrocliente);

-- Constraints tabla comercio
alter table comercio add constraint comercio_pk primary key (nrocomercio);

-- Constraints tabla compra
alter table compra add constraint compra_pk primary key (nrooperacion);
alter table compra add constraint compra_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
alter table compra add constraint compra_nrocomercio_fk foreign key (nrocomercio) references comercio(nrocomercio);

-- Constraints tabla rechazo
alter table rechazo add constraint rechazo_pk primary key (nrorechazo);

alter table rechazo add constraint rechazo_nrocomercio_fk foreign key (nrocomercio) references comercio (nrocomercio);

-- Constraints tabla cierre  
alter table cierre add constraint cierre_pk primary key (anio, mes, terminacion);

-- Constraints tabla cabecera
alter table cabecera add constraint cabecera_pk primary key (nroresumen);
alter table cabecera add constraint cabecera_nrotarjeta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);

-- Constraints tabla detalle 
alter table detalle add constraint detalle_pk primary key (nroresumen, nrolinea);
alter table detalle add constraint detalle_nroresumen_fk foreign key (nroresumen) references cabecera (nroresumen);

-- Constraints tabla alerta
alter table alerta add constraint alerta_pk primary key (nroalerta);

