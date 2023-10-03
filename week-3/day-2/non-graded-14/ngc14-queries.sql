-- Challenge 1: Product Count by Price Category
WITH Categories AS (
    SELECT
        CASE
            WHEN Price < 300 THEN 'Low Price'
            WHEN Price >= 300 AND Price <= 600 THEN 'Medium Price'
            ELSE 'High Price'
        END AS Category
    FROM Products
)
SELECT Category, COUNT(Category) AS Counter
FROM Categories
GROUP BY Category;

-- Challenge 2: Top N Products by Sales
SELECT p.ProductName, ps.QuantitySold
FROM ProductSales ps
JOIN Products p ON p.ProductID = ps.ProductID
ORDER BY ps.QuantitySold DESC
LIMIT 5;

-- Challenge 3: Product Sales Growth
WITH CurrentPeriod AS (
    SELECT ProductID, SUM(QuantitySold) AS CurrentSales
    FROM ProductSales
    WHERE
        SaleDate BETWEEN DATE_SUB(CURRENT_DATE, INTERVAL 1 MONTH) AND CURRENT_DATE
    GROUP BY ProductID
),
PreviousPeriod AS (
    SELECT ProductID, SUM(QuantitySold) AS PreviousSales
    FROM ProductSales
    WHERE
        SaleDate BETWEEN DATE_SUB(CURRENT_DATE, INTERVAL 2 MONTH) AND DATE_SUB(CURRENT_DATE, INTERVAL 1 MONTH) 
    GROUP BY ProductID
)
SELECT 
    c.ProductID,
    ps.ProductName,
    c.CurrentSales,
    COALESCE(P.PreviousSales, 0) AS PreviousSales,
    ROUND(
        IFNULL(((c.CurrentSales - p.PreviousSales) / NULLIF(p.PreviousSales, 0)), 1) * 100, 2
    ) AS PercentageGrowth
FROM CurrentPeriod c
LEFT JOIN PreviousPeriod p ON c.ProductID = p.ProductID
JOIN Products ps ON c.ProductID = ps.ProductID;


-- Challenge 4: Average Price Difference
WITH Averages AS (
    SELECT AVG(Price) AS AveragePrice
	FROM Products
)
SELECT p.ProductID, p.ProductName, p.Price - a.AveragePrice AS AVGPriceDifference
FROM Products p, Averages a
WHERE p.ProductID != 1 AND p.Price IS NOT NULL;


-- Challenge 5: Products without Price
SELECT ProductID, ProductName
FROM Products
WHERE Price IS NULL;