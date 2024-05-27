    REM Seteamos solo para este script que genere un ejecutale para OS linux
    set GOOS=linux

    REM Seteamos solo para este script que genere un ejecutale para arquitectura amd64
    set GOARCH=amd64

    REM Generamos el ejecutale
    REM go build main.go
    go build -tags lambda.norpc -o bootstrap main.go

    REM ******************************************************************************

    REM Amazon pide que las lambdas se suban Zipeadas

    REM Elimina main.zip en caso de existir, en caso contrario nos indica que no se encuentra el archivo.
    del main.zip

    REM Creamos archivo zip (zipeamos)
    tar.exe -a -cf main.zip bootstrap