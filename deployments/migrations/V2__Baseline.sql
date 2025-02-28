

CREATE TABLE IF NOT EXISTS `Psp` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `PspCode` varchar(16) NOT NULL,
  `PspShortName` varchar(50) DEFAULT NULL,
  `PspCountryCode` varchar(3) DEFAULT NULL,
  `CreatedAt` datetime NOT NULL,
  `UpdatedAt` datetime NOT NULL,
  `DeletedAt` datetime DEFAULT NULL,
  `CreatedBy` BINARY(16) NOT NULL, 
  `UpdatedBy` BINARY(16) DEFAULT NULL,
  `DeletedBy` BINARY(16) DEFAULT NULL,
  PRIMARY KEY (`ID`)
);




CREATE TABLE IF NOT EXISTS `FeeSet` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `PspID` int DEFAULT NULL,
  `IsActive` BOOL DEFAULT TRUE,
  `CreatedAt` datetime NOT NULL,
  `UpdatedAt` datetime DEFAULT NULL,
  `DeletedAt` datetime DEFAULT NULL,
  `CreatedBy` BINARY(16),
  `UpdatedBy` BINARY(16),
  `DeletedBy` BINARY(16) DEFAULT NULL,

  PRIMARY KEY (`ID`),


  CONSTRAINT `FeeSet_Psp_FK` FOREIGN KEY (`PspID`)
    REFERENCES `Psp` (`ID`)
    ON DELETE SET NULL

);


CREATE TABLE IF NOT EXISTS `FeeRange` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `FeeSetID` int DEFAULT NULL,
  `From` DECIMAL(18, 2) NOT NULL,
  `To`  DECIMAL(18, 2) NULL,
  `FeeFixed` DECIMAL(18, 2) NOT NULL,
  `FeePercentage` DECIMAL(9, 6) NOT NULL,
  `MaxTotalFee` DECIMAL(14, 2),
  `CreatedAt` datetime NOT NULL,
  `UpdatedAt` datetime DEFAULT NULL,
  `DeletedAt` datetime DEFAULT NULL,
  `CreatedBy` BINARY(16),
  `UpdatedBy` BINARY(16),
  `DeletedBy` BINARY(16) DEFAULT NULL,

  PRIMARY KEY (`ID`),

  CONSTRAINT `FeeRange_FeeSet_FK` FOREIGN KEY (`FeeSetID`)
    REFERENCES `FeeSet` (`ID`)
    ON DELETE SET NULL

);