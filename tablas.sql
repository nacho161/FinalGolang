CREATE TABLE Registro(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    IDActividad int,
    HoraInicio VARCHAR(12),
	HoraFinal VARCHAR(12),
    IDEmpleado int,
    Fecha date
);

CREATE TABLE Actividad(
    IDActividad INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    Descripcion VARCHAR(MAX),
    Proyecto VARCHAR(MAX)
);

CREATE TABLE Empleado(
    IDEmpleado INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    Nombre VARCHAR(MAX),
    Edad int,
	Cargo VARCHAR(MAX)
);