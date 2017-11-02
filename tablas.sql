CREATE TABLE `tiempos`.`registro` (
  
	`idregistro` INT NOT NULL AUTO_INCREMENT,
  
	`IDActividad` INT NOT NULL,
  
	`HoraInicio` VARCHAR(45) NULL,
  
	`IDEmpleado` INT NULL,
  
	`Fecha` DATETIME NULL,
  
PRIMARY KEY (`idregistro`));

CREATE TABLE `tiempos`.`actividad` (
  
	`idActividad` INT NOT NULL AUTO_INCREMENT,
  
	`Descripcion` VARCHAR(45) NULL,
  
	`Proyecto` VARCHAR(45) NULL,
  
PRIMARY KEY (`idActividad`));

CREATE TABLE `tiempos`.`empleado` (
  
	`idEmpleado` INT NOT NULL AUTO_INCREMENT,
  
	`nombre` VARCHAR(45) NULL,
  `edad` INT NULL,
  
	`Empleadocol` VARCHAR(45) NULL,
  
PRIMARY KEY (`idEmpleado`));