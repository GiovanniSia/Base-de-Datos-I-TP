CREATE OR REPLACE FUNCTION fn_alerta_clientes() 
RETURNS TRIGGER AS $fn_alerta_clientes$
  DECLARE
  
  v_fecha_ultima timestamp;
  
  BEGIN
     
    ----Si una tarjeta registra dos compras en un lapso menor de un minuto en comercios
    ----distintos ubicados en el mismo código postal	 
    select co.fecha
	into v_fecha_ultima
	from compra co,
	     comercio comp
	where co.nrocomercio = comp.nrocomercio
	and co.nrotarjeta = new.nrotarjeta
	AND co.nrocomercio <>  new.nrocomercio
	AND comp.codigopostal = (SELECT codigopostal 
							 FROM comercio WHERE nrocomercio = new.nrocomercio)
	order by nrooperacion desc
	limit 1;



	if (v_fecha_ultima is not null 
	and v_fecha_ultima > CURRENT_TIMESTAMP - (1 * interval '1 minute')) 
	then


	--Inserto Alerta por 1 minuto
			INSERT INTO alerta
			VALUES 
			(
			nextval('seq_nroalerta') , 
			new.nrotarjeta, 
			CURRENT_DATE,
			null, --ya que no hay rechazo relacionado
			1,
			'Compra en menos de 1 min'
			); 
	
	else
	----Si una tarjeta registra dos compras en un lapso menor de 5 minutos en comercios
	----con diferentes códigos postales.
	    	select fecha
			into v_fecha_ultima
			from compra co,
				 comercio comp
			where co.nrocomercio = comp.nrocomercio
			and co.nrotarjeta = new.nrotarjeta
			AND co.nrotarjeta = new.nrotarjeta
			AND co.nrocomercio <>  new.nrocomercio
			AND comp.codigopostal <> (SELECT codigopostal 
									 FROM comercio WHERE nrocomercio = new.nrocomercio)
			order by nrooperacion desc
			limit 1;



			if (v_fecha_ultima is not null 
			and v_fecha_ultima > CURRENT_TIMESTAMP - (1 * interval '5 minute')) 
			then


			--Inserto Alerta por 5 minutoa
					INSERT INTO alerta
					VALUES 
					(
					nextval('seq_nroalerta') , 
					new.nrotarjeta, 
					CURRENT_DATE,
					null, --ya que no hay rechazo relacionado
					5,
					'Compra en menos de 5 min'
					); 
					
			end if;
	

	end if;




   RETURN NEW;
  END;
$fn_alerta_clientes$ 
LANGUAGE plpgsql;
