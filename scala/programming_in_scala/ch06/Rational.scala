import java.io.File

object test {
	def scalaFiles = 
		for {
			file <- filesHere
			if file.getName.endsWith(".scala")
		} yield file

	def main(args: Array[String]) {
	}
}
