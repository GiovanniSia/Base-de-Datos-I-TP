     
-------------------------------------------------------------------------------
--AUTORIZACION_COMPRA
-------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION autorizacion_compra(numero_tarjeta char, codigo_seguridad char, numero_comercio int, el_monto decimal)
  RETURNS boolean AS
$BODY$
DECLARE

	v_nrotarjeta char(16);
	v_codseguridad char(4);
	v_estado character(10);
	v_limitecompra decimal(8,2);
	v_validahasta char(6);
	
	v_montopendiente decimal(7,2);
	
	v_mesactual char(2);
	v_anioactual char(4);
	
	v_cant_rechazos_lim  int;
	
BEGIN
	
	
	SELECT p.nrotarjeta, p.codseguridad, p.estado, p.limitecompra, p.validahasta
	INTO v_nrotarjeta, v_codseguridad, v_estado, v_limitecompra, v_validahasta
	FROM tarjeta p
	WHERE p.nrotarjeta = numero_tarjeta ;

	-- Caso 1a, la tarjeta no existe	
	IF (v_nrotarjeta IS NULL)
	THEN
		
	    --Inserto Rechazo
		INSERT INTO rechazo
		VALUES 
		(
		nextval('seq_nrorechazo') , 
		numero_tarjeta,
		numero_comercio,
		CURRENT_DATE,
		el_monto, 
		'tarjeta no valida o no vigente'
		); 

		--Inserto Alerta
		INSERT INTO alerta
		VALUES 
		(
		nextval('seq_nroalerta') , 
		numero_tarjeta, 
		CURRENT_DATE,
		currval('seq_nrorechazo') , 
		0,
		'tarjeta no valida o no vigente'
		); 

		
		RETURN false;
		
	END IF;
	
	
	-- Caso 5, la tarjeta esta suspendida , lo pongo aca para que lo pregunte antes del siguiente, si no nunca saltaria
	IF (v_estado = 'suspendida')
	THEN
	
	    --Inserto Rechazo
		INSERT INTO rechazo
		VALUES 
		(
		nextval('seq_nrorechazo'), 
		numero_tarjeta,
		numero_comercio,
		CURRENT_DATE,
		el_monto, 
		'la tarjeta se encuentra suspendida'
		); 
		
		
		 --Inserto Alerta
		INSERT INTO alerta
		VALUES 
		(
		nextval('seq_nroalerta') , 
		numero_tarjeta, 
		CURRENT_DATE,
		currval('seq_nrorechazo') , 
		0,
		'la tarjeta se encuentra suspendida'
		); 
		
		RETURN false;
		
	END IF;
	
	
	-- Caso 1b, la tarjeta no esta vigente
	IF (v_estado <> 'vigente')
	THEN
	
	    --Inserto Rechazo
		INSERT INTO rechazo
		VALUES 
		(
		nextval('seq_nrorechazo'), 
		numero_tarjeta,
		numero_comercio,
		CURRENT_DATE,
		el_monto, 
		'tarjeta no valida o no vigente'
		); 
		
		
			--Inserto Alerta
		INSERT INTO alerta
		VALUES 
		(
		nextval('seq_nroalerta') , 
		numero_tarjeta, 
		CURRENT_DATE,
		currval('seq_nrorechazo') , 
		0,
		'tarjeta no valida o no vigente'
		); 
		RETURN false;
		
	END IF;
	
	--Caso 2, el codigo de seguridad es invalido
	IF (v_codseguridad <> codigo_seguridad)
	THEN
			
		--Inserto Rechazo
		INSERT INTO rechazo
		VALUES 
		(
		nextval('seq_nrorechazo'), 
		numero_tarjeta,
		numero_comercio,
		CURRENT_DATE,
		el_monto, 
		'codigo de seguridad invalido'
		); 
		
	    --Inserto Alerta
		INSERT INTO alerta
		VALUES 
		(
		nextval('seq_nroalerta') , 
		numero_tarjeta, 
		CURRENT_DATE,
		currval('seq_nrorechazo') , 
		0,
		'codigo de seguridad invalido'
		); 

		
		RETURN false;
		
	END IF;

	
	--Caso 3, supera el limite de compra
	
	--Sumo el total de sus compras
    SELECT SUM(monto) 
	into v_montopendiente 
	FROM compra c
	WHERE v_nrotarjeta=c.nrotarjeta
	AND c.pagado = false; 
	

	v_montopendiente := v_montopendiente + el_monto;


	IF (v_montopendiente > v_limitecompra)
	THEN
	    --Inserto Rechazo
		INSERT INTO rechazo
		VALUES 
		(
		nextval('seq_nrorechazo'), 
		numero_tarjeta,
		numero_comercio,
		CURRENT_DATE,
		el_monto, 
		'supera limite de tarjeta'
		); 
		
		 --Inserto Alerta
		INSERT INTO alerta
		VALUES 
		(
		nextval('seq_nroalerta') , 
		numero_tarjeta, 
		CURRENT_DATE,
		currval('seq_nrorechazo') , 
		0,
		'supera limite de tarjeta'
		); 
		
		
		select count(nrotarjeta) 
		into v_cant_rechazos_lim
		from rechazo where nrotarjeta = numero_tarjeta 
		and motivo ='supera limite de tarjeta'
		and fecha = CURRENT_DATE ;
		
		
		if (v_cant_rechazos_lim >1)then
		
		   update tarjeta set estado='suspendida' where nrotarjeta=numero_tarjeta;
		   
		   
		    --Inserto Alerta
    		INSERT INTO alerta
    		VALUES 
    		(
    		nextval('seq_nroalerta') , 
    		numero_tarjeta, 
    		CURRENT_DATE,
    		currval('seq_nrorechazo') , 
    		0,
    		'suspencion preventiva'
    		); 
		
		end if;
		
		
		
		RETURN false;
	END IF;

	--Caso 4, la tarjeta esta vencida
	
	SELECT EXTRACT(Year FROM CURRENT_DATE) into v_anioactual;
	SELECT EXTRACT(Month FROM CURRENT_DATE) into v_mesactual;
	
	IF (v_anioactual|| v_mesactual > v_validahasta)
	THEN
	
	    --Inserto Rechazo
		INSERT INTO rechazo
		VALUES 
		(
		nextval('seq_nrorechazo'), 
		numero_tarjeta,
		numero_comercio,
		CURRENT_DATE,
		el_monto, 
		'plazo de vigencia expirado'
		); 
		
		 --Inserto Alerta
		INSERT INTO alerta
		VALUES 
		(
		nextval('seq_nroalerta') , 
		numero_tarjeta, 
		CURRENT_DATE,
		currval('seq_nrorechazo') , 
		0,
		'plazo de vigencia expirado'
		); 
		
		RETURN false;
		
	END IF;
	
	
	
    
------------------------------------------------	
	
	
	--Si todo esta bien
	INSERT INTO compra
		VALUES 
		(
		nextval('seq_nrocompra'), 
		numero_tarjeta,
		numero_comercio,
		CURRENT_DATE,
		el_monto, 
		false
		); 
		
		RETURN true;
	
	
END;

$BODY$
  LANGUAGE plpgsql
