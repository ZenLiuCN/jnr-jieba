package cn.zenliu.kotlin.jieba

import cn.zenliu.kotlin.jieba.JnrJieba.TokenizerMode.CUT_ALL
import jnr.ffi.LibraryLoader
import jnr.ffi.annotations.Encoding
import jnr.ffi.annotations.Transient
import java.io.File
import java.io.InputStream
import java.security.MessageDigest
import javax.xml.bind.DatatypeConverter

object JnrJieba {
	private const val soName = "jieba.so"
	private const val dictName = "dict.gz"
	private const val dictNameBig = "dict.big.gz"
	private val InputStream.md5sum
		get() = MessageDigest.getInstance("MD5").let { md5 ->
			buffered().use {
				md5.digest(it.readBytes())
			}.let {
				DatatypeConverter.printHexBinary(it)
			}
		}
	private val jr = loadLibFromJar()
	private var initialization = false
	private fun loadLibFromJar(): Jieba {
		val soo = File(soName)
		val dico = File(dictName)
		val dicbo = File(dictNameBig)

		val so = this::class.java.getResourceAsStream("/jnr/$soName")
		val dict = this::class.java.getResourceAsStream("/jnr/$dictName")
		val dictBig = this::class.java.getResourceAsStream("/jnr/$dictNameBig")

		if (soo.exists() && dico.exists() && dicbo.exists()) {
			//control by user
		/*	if (
					soo.inputStream().md5sum == so.md5sum
					&& dico.inputStream().md5sum == dict.md5sum
					&& dicbo.inputStream().md5sum == dictBig.md5sum
			)*/
				return LibraryLoader.create(Jieba::class.java).load(File(soName).absolutePath)
		}

		soo.apply {
			createNewFile()
			outputStream().use {
				so.copyTo(it)
			}
		}
		dico.apply {
			createNewFile()
			outputStream().use {
				dict.copyTo(it)
			}
		}
		dicbo.apply {
			createNewFile()
			outputStream().use {
				dictBig.copyTo(it)
			}
		}
		return LibraryLoader.create(Jieba::class.java).load(File(soName).absolutePath)
	}

	internal interface Jieba {
		fun AddWord(@Transient @Encoding("UTF8") word: String, frequency: Double)
		fun RemoveWord(@Transient @Encoding("UTF8") word: String)
		fun Initial(@Transient @Encoding("UTF8") dictPath: String)

		@Encoding("UTF8")
		fun Tokenizer(@Transient @Encoding("UTF8") src: String, @Encoding("UTF8") join: String, mode: Int): String
	}

	/**
	 *
	 * @param dictionaryPath String default is dict.gz
	 * @return Unit
	 */
	fun initial(dictionaryPath: String = dictName) {
		if (initialization) return
		jr.Initial(dictionaryPath)
		initialization = true
	}

	/**
	 *
	 * @param src String
	 * @param delimiter String
	 * @param mode TokenizerMode
	 * @return List<String>
	 */
	fun tokenizer(src: String, delimiter: String = "|", mode: TokenizerMode = CUT_ALL): List<String> {
		assert(initialization) { "jieba not be initialized" }
		assert(delimiter.length in 1..3) { "delimiter must have length of 1-3" }
		if (src.isBlank()) return emptyList()
		return jr.Tokenizer(src, delimiter, mode.ordinal).split(delimiter)
	}

	/**
	 * temple add word to dictionary
	 * @param word String
	 * @param frequency Double
	 * @return Unit
	 */
	fun addWord(word: String, frequency: Double) {
		assert(initialization) { "jieba not be initialized" }
		jr.AddWord(word, frequency)
	}

	/**
	 * remove a word from dictionary
	 * @param word String
	 * @return Unit
	 */
	fun removeWord(word: String) {
		assert(initialization) { "jieba not be initialized" }
		jr.RemoveWord(word)
	}

	enum class TokenizerMode {
		CUT_ALL,
		CUT_HMM,
		CUT_SEARCH_HMM,
		CUT_SEARCH_NO_HMM,
		CUT_NO_HMM,
		;
	}
}