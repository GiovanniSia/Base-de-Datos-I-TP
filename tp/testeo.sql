-------------------------------------------------------------------------------
--Procedimiento de Testeo
-------------------------------------------------------------------------------


CREATE OR REPLACE FUNCTION prc_testeo()
RETURNS void AS
$BODY$
DECLARE

	prueba_datos CURSOR 
	FOR SELECT nrotarjeta, codseguridad, nrocomercio, monto 
	FROM consumo;
				
	v_nrotarjeta char(16);
	v_codseguridad char(4);
	v_nrocomercio int;
	v_monto decimal(7,2);
		
BEGIN
	
	OPEN prueba_datos;
	LOOP 
	FETCH prueba_datos	
	INTO v_nrotarjeta, v_codseguridad, v_nrocomercio, v_monto;

	IF NOT FOUND THEN
		EXIT;
	END IF;
	
		PERFORM "autorizacion_compra"
		(
		v_nrotarjeta,
		v_codseguridad,
		v_nrocomercio, 
		v_monto
		);

	END LOOP;
	
	CLOSE prueba_datos;

	RETURN;

	
END;
$BODY$
LANGUAGE plpgsql;
  
