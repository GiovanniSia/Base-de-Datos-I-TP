create or replace function funcionresumen(numcliente int, periodo int) returns void as $$  
  declare 
        termtarjeta int; 
        numresumen int;
        total decimal(7,2);
        numlinea int; 
        mes int;
        nomcomercio text;
        elcliente record;
        sutarjeta record;
        sucierre record;
        v record;
        w record;
     
  begin  
        numresumen:= 1;
        numresumen:= numresumen + count(*) from cabecera; 
        total:= 0; 
        numlinea:= 1; 
         
        select * into elcliente from cliente where nrocliente = numcliente; 
        select * into sutarjeta from tarjeta where nrocliente = numcliente; 
        termtarjeta:= substr(sutarjeta.nrotarjeta,16); 
        select * into sucierre from cierre where terminacion = termtarjeta; 
        
        for v in select * from compra loop
          mes:= extract(month from v.fecha);
          if mes = periodo and v.nrotarjeta = sutarjeta.nrotarjeta then 
             total:= total + v.monto;
          end if;  
        end loop;
        
        insert into cabecera values(numresumen,elcliente.nombre,elcliente.apellido,elcliente.domicilio,sutarjeta.nrotarjeta,sucierre.fechainicio,sucierre.fechacierre,sucierre.fechavto,total);
          
        for w in select * from compra loop  
          if w.nrotarjeta = sutarjeta.nrotarjeta and mes = periodo then
              select nombre into nomcomercio from comercio where nrocomercio = w.nrocomercio; 
             insert into detalle values (numresumen,numlinea,w.fecha,nomcomercio,w.monto);
             numlinea:= numlinea +1;
          end if;
        end loop;  
        numlinea:= 1;  
  end;
$$ language plpgsql;
