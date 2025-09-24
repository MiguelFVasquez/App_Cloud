@echo off
REM compilación cruzada para Linux (amd64) desde Windows

setlocal

REM Cambia si quieres otra arquitectura (arm64, 386, etc.)
set GOOS=linux
set GOARCH=amd64

REM Desactivar CGO para builds totalmente estáticos/puros en Go
set CGO_ENABLED=0

echo ============================================
echo Compilando para %GOOS% / %GOARCH% (CGO=%CGO_ENABLED%)
echo ============================================

REM Construye el binario desde la ruta del main; ajusta la ruta si tu main no está en ./cmd/server
go build -o srvimglinux ./cmd/server

if %ERRORLEVEL% neq 0 (
  echo.
  echo *** ERROR: La compilación falló.
  endlocal
  exit /b %ERRORLEVEL%
)

echo.
echo Compilación completada: srvimglinux
echo El binario está listo para copiar a un sistema Linux.
endlocal
