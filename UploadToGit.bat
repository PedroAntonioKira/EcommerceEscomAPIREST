REM ESTE ES UN COMENTARIO QUE NO APARECE EN CONSOLA PARA WINDOWS

REM VISUALIZAMOS QUE ARCHIVOS HAN SIDO CREADOS O MODIFICADOS Y AUN NO SE HAN SUBIDO A GIT
REM LOS QUE NO SE HAN SUBIDO QUE HAN SIDO CREADOS O MODIFICADOS APARECERAN EN ROJO
REM LOS QUE YA ESTAN SE SUBIERON AL REPOSITORIO LOCAL APARECEN EN VERDE
git status

REM SE AGREGA TODOS LOS ARCHIVOS QUE SE HAN MODIFICADO Y SE SUBEN AL GIT
git add .

REM SE CONFIGURAN CREDENCIALES PARA SUBIR
REM SE CONFIGURA EL USERMAIL ASOCIADO AL GIT QUE SE VA A SUBIR
git config --global user.email "morales.gonzalez.pedroantonio@gmail.com"
REM SE CONFIGURA EL USERNAME ASOCIADO AL GIT QUE SE VA A SUBIR
git config --global user.name "PedroAntonioKira"

REM SE AGREGA UN COMENTARIO
git commit -m "Primera Prueba Completa para conexión con el resto del backend y apigateway parte 02 Products"

REM SUBE TODO EL PROYECTO A GITHUB
git push