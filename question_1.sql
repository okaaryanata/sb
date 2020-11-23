SELECT u.ID, u.UserName, 
 (SELECT UserName FROM USER WHERE ID=u.Parent) AS ParentUserName
FROM USER AS u;