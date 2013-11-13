<html>
	<head>
		<title>Sentencegator - web interface</title>
	</head>
	<body background="background.jpg">
		<form method="POST" action="sentences">
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
					<td>Include B-lines</td>
					<td><input type="checkbox" name="bl" value="true" /></td>
				</tr>
				<tr><td><input type="submit" value="Get Sentences" /></td>
				</tr>
			</table>
		</form>
	</body>
</html>
