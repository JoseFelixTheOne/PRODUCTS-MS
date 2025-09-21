-- Crear base de datos si no existe
IF DB_ID('ProductsDB') IS NULL
    CREATE DATABASE ProductsDB;
GO

USE ProductsDB;
GO

-- Crear tabla Category
IF OBJECT_ID('dbo.Category', 'U') IS NULL
BEGIN
    CREATE TABLE dbo.Category (
        CategoryID INT IDENTITY(1,1) PRIMARY KEY,
        Name NVARCHAR(100) NOT NULL UNIQUE,
        Slug NVARCHAR(120) NOT NULL UNIQUE,
        CreatedAt DATETIME2 NOT NULL DEFAULT SYSUTCDATETIME(),
        UpdatedAt DATETIME2 NOT NULL DEFAULT SYSUTCDATETIME()
    );
END
GO

-- Crear tabla Product
IF OBJECT_ID('dbo.Product', 'U') IS NULL
BEGIN
    CREATE TABLE dbo.Product (
        ProductID INT IDENTITY(1,1) PRIMARY KEY,
        SKU NVARCHAR(64) NOT NULL UNIQUE,
        Name NVARCHAR(200) NOT NULL,
        Description NVARCHAR(MAX) NULL,
        Price DECIMAL(18,2) NOT NULL,
        Stock INT NOT NULL DEFAULT 0,
        Active BIT NOT NULL DEFAULT 1,
        CategoryID INT NOT NULL,
        CreatedAt DATETIME2 NOT NULL DEFAULT SYSUTCDATETIME(),
        UpdatedAt DATETIME2 NOT NULL DEFAULT SYSUTCDATETIME(),
        CONSTRAINT FK_Product_Category FOREIGN KEY (CategoryID) REFERENCES dbo.Category(CategoryID)
    );

    CREATE INDEX IX_Product_Name ON dbo.Product(Name);
    CREATE INDEX IX_Product_Price ON dbo.Product(Price);
    CREATE INDEX IX_Product_Stock ON dbo.Product(Stock);
    CREATE INDEX IX_Product_Active ON dbo.Product(Active);
    CREATE INDEX IX_Product_CategoryID ON dbo.Product(CategoryID);
END
GO

-- Insertar categorías si no existen
IF NOT EXISTS (SELECT 1 FROM dbo.Category)
BEGIN
    INSERT INTO dbo.Category (Name, Slug)
    VALUES 
        ('Hogar','hogar'), 
        ('Ropa','ropa'),
        ('Comida Vegana','comida-vegana');
END
GO

-- Insertar productos de ejemplo
IF NOT EXISTS (SELECT 1 FROM dbo.Product)
BEGIN
    INSERT INTO dbo.Product (SKU, Name, Description, Price, Stock, Active, CategoryID)
    VALUES
        ('FOOD-001','Pan de Higo','Pan artesanal con higos secos', 15.50, 25, 1, 4),
        ('FOOD-002','Pan Integral de Higo','Pan integral enriquecido con higos', 17.00, 30, 1, 4),
        ('FOOD-003','Pan de Centeno','Pan rústico de centeno', 12.00, 15, 1, 4),
        ('FOOD-004','Leche de Almendra','Bebida vegetal sin lactosa', 8.90, 40, 1, 4),
        ('FOOD-005','Tofu','Proteína vegetal a base de soya', 10.50, 35, 1, 4);
END
GO
