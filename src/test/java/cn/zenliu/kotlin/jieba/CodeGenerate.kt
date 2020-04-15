package cn.zenliu.kotlin.jieba

import java.io.*

fun createTuples() {
	val f = File("out.txt").apply { createNewFile() }
	val i = File::class.java.getResource("/GenerateData.txt").file.let(::File)
	val keys = mutableSetOf<String>()
	i.forEachLine {
		val (k, v) = it.split("|")
		if (k.isNotBlank() && keys.add(k)) {
			f.appendText("`$k`:&PosTuple")
			f.appendText(v)
			f.appendText(",\n")
		}
	}

}

fun joinNames() {
	val f = "char_state_table.go".let(::File).apply { createNewFile() }
	val i = File::class.java.getResource("/char_state_tab.go").file.let(::File)

	i.forEachLine {
		if (it.indexOf("`") < 0) {
			f.appendText(it)
			f.appendText("\n")
		} else {
			f.appendText(it
				.replace("`", "")
				.replace("{'", "tuples[`")
				.replace("', ", "")
				.replace("}", "`]"))
			f.appendText("\n")
		}
	}

}

fun main() {
	joinNames()
}
