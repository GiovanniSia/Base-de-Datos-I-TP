package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
    "log"
    "io/ioutil"
    "strings"
	"encoding/json"
    bolt "github.com/coreos/bbolt"
    "strconv"
    "time"
)

func verificarError(err error){
		if err != nil {
        log.Fatal(err)
    }
}

func crearBaseDeDatos() {
    db,err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
    verificarError(err)
    defer db.Close()

    _, err = db.Exec(`create database tarjetas`)
    verificarError(err)
    
}

func crearTablas(){

	   db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()

    archivo, err :=ioutil.ReadFile("./crearTablas.sql") 
    verificarError(err)

    solicitud :=strings.Split(string(archivo),"\n")
    for _,peticion := range solicitud{
		db.Exec(peticion)
	}
}

func crearPKsFKs(){

	   db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()

    archivo, err :=ioutil.ReadFile("./crearPKsFKs.sql") 
    verificarError(err)

    solicitud :=strings.Split(string(archivo),"\n")
    for _,peticion := range solicitud{
		db.Exec(peticion)
	}	
}

func crearDatos(){	

		   db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()

    archivo, err :=ioutil.ReadFile("./crearDatos.sql") 
    verificarError(err)

    solicitud :=strings.Split(string(archivo),"\n")
    for _,peticion := range solicitud{
		db.Exec(peticion)
	}

	archivo, err = ioutil.ReadFile("./insertarCierre.sql")
	verificarError(err)
	_, err = db.Exec(string(archivo))
	verificarError(err)
	_, err = db.Exec(`select insert_cierre()`)
    verificarError(err)
     
}

func borrarBaseDeDatos(){
	
		db,err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
    verificarError(err)
    defer db.Close()
   
    _, err = db.Exec(`drop database tarjetas`)
    verificarError(err)
	
}

func borrarTablas(){
	
		   db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()

    archivo, err :=ioutil.ReadFile("./borrarTablas.sql") 
    verificarError(err)

    solicitud :=strings.Split(string(archivo),"\n")
    for _,peticion := range solicitud{
		db.Exec(peticion)
	}
	
}

//Borro PKs y FKs
func borrarPKsFKs(){
	 db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()

    archivo, err :=ioutil.ReadFile("./borrarPKsFKs.sql") 
    verificarError(err)

    solicitud :=strings.Split(string(archivo),"\n")
    for _,peticion := range solicitud{
		db.Exec(peticion)
	}
}

func borrarDatos(){
	
		   db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()

    archivo, err :=ioutil.ReadFile("./borrarDatos.sql") 
    verificarError(err)

    solicitud :=strings.Split(string(archivo),"\n")
    for _,peticion := range solicitud{
		db.Exec(peticion)
	}
	
}

func alerta(){
		   db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()

    archivo, err :=ioutil.ReadFile("./alertas.sql") 
    verificarError(err)
	
	_, err = db.Exec(string(archivo))
	verificarError(err)
}

func autorizarCompra(){	
			db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()

    archivo, err :=ioutil.ReadFile("./autorizarCompra.sql") 
    verificarError(err)

	_, err = db.Exec(string(archivo))
	verificarError(err)
	
		_, err = db.Exec(`CREATE TRIGGER alertas before INSERT ON compra FOR EACH ROW EXECUTE PROCEDURE fn_alerta_clientes();`)
    verificarError(err)	
	
}

func pruebaConsumo(){
	alerta()
	autorizarCompra()	
			db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()

    archivo, err :=ioutil.ReadFile("./testeo.sql") 
    verificarError(err)
    
	_, err = db.Exec(string(archivo))
	verificarError(err)


	// Alerta de 1 min
		_, err = db.Exec(`insert into compra values (15,4372364962513947,11, '2020-06-21 19:00:00' ,400.0,false)`)
    verificarError(err)	
    	_, err = db.Exec(`insert into compra values (16,4372364962513947,12, '2020-06-22 19:00:01' ,400.0,false)`)
    verificarError(err)	

	//Alerta de 5 min
		_, err = db.Exec(`insert into compra values (17,4526712738911625,15, '2020-06-21 19:00:00' ,400.0,false)`)
    verificarError(err)	
    	_, err = db.Exec(`insert into compra values (18,4526712738911625,13, '2020-06-22 19:00:01' ,400.0,false)`)
    verificarError(err)	

	
		_, err = db.Exec(`select prc_testeo()`)
    verificarError(err)	
}

func generarResumen(){
			   db, err := sql.Open("postgres", "user=postgres host=localhost dbname=tarjetas sslmode=disable")
    verificarError(err)
    defer db.Close()
   
	    archivo, err :=ioutil.ReadFile("./generarResumen.sql") 
    verificarError(err)

	_, err = db.Exec(string(archivo))
	verificarError(err)
	
	_, err = db.Exec(`select funcionresumen(20 , 6)`)
    verificarError(err)
	
}

type Cliente struct{
	Nrocliente int
	Nombre string
	Apellido string
	Domicilio string
	Telefono string				
}

type Comercio struct{
	Nrocomercio int
	Nombre string
	Domicilio string
	Codigopostal string
	Telefono string
}

type Tarjeta struct{
	Nrotarjeta string
	Nrocliente int
	Validadesde string
	Validahasta string
	Codseguridad string
	Limitecompra float64
	Estado string
}

type Compra struct{
	Nrooperacion int
	Nrotarjeta string
	Nrocomercio int
	Fecha time.Time
	Monto float64
	Pagado bool
}

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
    
    // abre transacción de escritura
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

    err = b.Put(key, val)
    if err != nil {
        return err
    }

    // cierra transacción
    if err := tx.Commit(); err != nil {
        return err
    }

    return nil
}
		
func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
    var buf []byte

    // abre una transacción de lectura
    err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucketName))
        buf = b.Get(key)
        return nil
    })

    return buf, err
}

func baseDeDatosNoSQL() {
    db, err := bolt.Open("tarjetasBolt.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	cliente1 := Cliente{20,"Mariana","Gutierrez","Av. Peron 1023, San Miguel","011536729328"}
	cliente2 := Cliente{21,"Ricardo","Martinez","Av. Gaspar Campos 3002, Jose C. Paz","011572636271"}
	cliente3 := Cliente{22,"Alan","Bruso","Av. Peron 2003, San Miguel","011546732819"}

	comercio1 := Comercio{10 , "Ferreteria Fiore" , "Santa Fe 2601, CABA","1800", "011556667890"}
	comercio2 := Comercio{11 , "Supermercado Ahorro" , "Av Cordoba 2012, CABA","1445", "011533244532"}
	comercio3 := Comercio{12 , "Papelera Lumier" , "Cervantes 344, Sol y Verde","1445", "011521112436"}
	
	tarjeta1 := Tarjeta{"4372364962513947",cliente1.Nrocliente,"0815","0820","2514",20000.00,"vigente"}
	tarjeta2 := Tarjeta{"4526712738911625",cliente2.Nrocliente,"0417","0420","5632",20000.00,"vigente"}
	tarjeta3 := Tarjeta{"5666578953335678",cliente3.Nrocliente,"0415","0420","5643",30000.00,"suspendida"}

	compra1 := Compra{1,tarjeta1.Nrotarjeta , comercio1.Nrocomercio , time.Now(),343.0,true}
	compra2 := Compra{2,tarjeta2.Nrotarjeta , comercio2.Nrocomercio , time.Now(),300.0,true}
	compra3 := Compra{3,tarjeta3.Nrotarjeta , comercio3.Nrocomercio , time.Now(),500.0,false}
	
	dataCliente1, err := json.Marshal(cliente1)
	verificarError(err)	
	 CreateUpdate(db   , "cliente", []byte(strconv.Itoa(cliente1.Nrocliente)), dataCliente1 )
	     resultadoCliente1, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(cliente1.Nrocliente)))
    fmt.Printf("%s\n", resultadoCliente1)
    
    
	dataCliente2, err := json.Marshal(cliente2)
	verificarError(err)	
	 CreateUpdate(db   , "cliente", []byte(strconv.Itoa(cliente2.Nrocliente)), dataCliente2 )
	     resultadoCliente2, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(cliente2.Nrocliente)))
    fmt.Printf("%s\n", resultadoCliente2)
    
    
	dataCliente3, err := json.Marshal(cliente3)
	verificarError(err)	
	 CreateUpdate(db   , "cliente", []byte(strconv.Itoa(cliente3.Nrocliente)), dataCliente3 )
	     resultadoCliente3, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(cliente3.Nrocliente)))
    fmt.Printf("%s\n", resultadoCliente3)
    
    
	dataComercio1, err := json.Marshal(comercio1)
	verificarError(err)	
	 CreateUpdate(db   , "comercio", []byte(strconv.Itoa(comercio1.Nrocomercio)), dataComercio1 )
	     resultadoComercio1, err := ReadUnique(db, "comercio", []byte(strconv.Itoa(comercio1.Nrocomercio)))
    fmt.Printf("%s\n", resultadoComercio1)
    
    
	dataComercio2, err := json.Marshal(comercio2)
	verificarError(err)	
	 CreateUpdate(db   , "comercio", []byte(strconv.Itoa(comercio2.Nrocomercio)), dataComercio2 )
	     resultadoComercio2, err := ReadUnique(db, "comercio", []byte(strconv.Itoa(comercio2.Nrocomercio)))
    fmt.Printf("%s\n", resultadoComercio2)
    
    
	dataComercio3, err := json.Marshal(comercio3)
	verificarError(err)	
	 CreateUpdate(db   , "comercio", []byte(strconv.Itoa(comercio3.Nrocomercio)), dataComercio3 )
	     resultadoComercio3, err := ReadUnique(db, "comercio", []byte(strconv.Itoa(comercio3.Nrocomercio)))
    fmt.Printf("%s\n", resultadoComercio3)
    
    
	dataTarjeta1, err := json.Marshal(tarjeta1)
	verificarError(err)	
	 CreateUpdate(db   , "tarjeta", []byte(tarjeta1.Nrotarjeta), dataTarjeta1 )
	     resultadoTarjeta1, err := ReadUnique(db, "tarjeta", []byte(tarjeta1.Nrotarjeta))
    fmt.Printf("%s\n", resultadoTarjeta1)
    
    
	dataTarjeta2, err := json.Marshal(tarjeta2)
	verificarError(err)	
	 CreateUpdate(db   , "tarjeta", []byte(tarjeta2.Nrotarjeta), dataTarjeta2 )
	     resultadoTarjeta2, err := ReadUnique(db, "tarjeta", []byte(tarjeta2.Nrotarjeta))
    fmt.Printf("%s\n", resultadoTarjeta2)
    
    
	dataTarjeta3, err := json.Marshal(tarjeta3)
	verificarError(err)	
	 CreateUpdate(db   , "tarjeta", []byte(tarjeta3.Nrotarjeta), dataTarjeta3 )
	     resultadoTarjeta3, err := ReadUnique(db, "tarjeta", []byte(tarjeta3.Nrotarjeta))
    fmt.Printf("%s\n", resultadoTarjeta3)
    
    
	dataCompra1, err := json.Marshal(compra1)
	verificarError(err)	
	 CreateUpdate(db   , "compra", []byte(strconv.Itoa(compra1.Nrooperacion)), dataCompra1 )
	     resultadoCompra1, err := ReadUnique(db, "compra", []byte(strconv.Itoa(compra1.Nrooperacion)))
    fmt.Printf("%s\n", resultadoCompra1)
    
    
	dataCompra2, err := json.Marshal(compra2)
	verificarError(err)	
	 CreateUpdate(db   , "compra", []byte(strconv.Itoa(compra2.Nrooperacion)), dataCompra2 )
	     resultadoCompra2, err := ReadUnique(db, "compra", []byte(strconv.Itoa(compra2.Nrooperacion)))
    fmt.Printf("%s\n", resultadoCompra2)
    
    
	dataCompra3, err := json.Marshal(compra3)
	verificarError(err)	
	 CreateUpdate(db   , "compra", []byte(strconv.Itoa(compra3.Nrooperacion)), dataCompra3 )
	     resultadoCompra3, err := ReadUnique(db, "compra", []byte(strconv.Itoa(compra3.Nrooperacion)))
    fmt.Printf("%s\n", resultadoCompra3)	
}

func imprimirMenu(){
		fmt.Print("0. Salir\n")
		fmt.Print("-------------------------------------------\n")
		fmt.Print("1. Crear base de datos\n")
		fmt.Print("-------------------------------------------\n")
		fmt.Print("2. Crear tablas\n")
		fmt.Print("-------------------------------------------\n")
        fmt.Print("3. Crear PKs y FKs \n")
        fmt.Print("-------------------------------------------\n")
        fmt.Print("4. Crear datos\n")
        fmt.Print("-------------------------------------------\n")
		fmt.Print("5. Borrar base de datos\n")
		fmt.Print("-------------------------------------------\n")
        fmt.Print("6. Borrar PKs y FKs\n")
        fmt.Print("-------------------------------------------\n")
        fmt.Print("7. Borrar datos\n")
        fmt.Print("-------------------------------------------\n")
        fmt.Print("8. Borrar tablas\n")
        fmt.Print("-------------------------------------------\n")
        fmt.Print("9. Prueba Consumo\n")
        fmt.Print("-------------------------------------------\n")
        fmt.Print("10. Generar resumen \n")
        fmt.Print("-------------------------------------------\n")
        fmt.Print("11. Crear y mostrar base de datos No-SQL\n")
        fmt.Print("-------------------------------------------\n")
}

func main() {
	menu:=true
	for menu{		
		imprimirMenu()
		
		var eleccion int	
		fmt.Print("Ingrese un numero: ")
	    fmt.Scanf("%d", &eleccion)
		
		switch eleccion {
		case 0:
			menu=false
		case 1:
			crearBaseDeDatos()
		case 2: 
			crearTablas()
		case 3:
			crearPKsFKs()
		case 4:
			crearDatos()
		case 5:
			borrarBaseDeDatos()
		case 6:
			borrarPKsFKs()
		case 7:
			borrarDatos()
		case 8:
			borrarTablas()
		case 9:
			pruebaConsumo()
		case 10:
			generarResumen()
		case 11:
			baseDeDatosNoSQL()
		default:
			fmt.Println("Eleccion erronea")
		}
	}
}
