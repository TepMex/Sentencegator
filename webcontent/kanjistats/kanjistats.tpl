<html>
	<head>
		<title>KanjiStats - web interface</title>
	</head>
	<body background="background.jpg">
		<form method="POST" action="stats"  enctype="multipart/form-data">
			<table>
				<tr>
					<td>API Key</td>
					<td><input type="text" name="apik" /></td>
				</tr>
				<tr>
					<td>Levels</td>
					<td><input type="text" name="levels" /></td>
				</tr>
				<tr>
					<td>File</td>
					<td><input type="file" name="japfile"/></td>
				</tr>
				<tr><td><input type="submit" value="Get Statistics" /></td>
				</tr>
			</table>
		</form>
	</body>
</html>
