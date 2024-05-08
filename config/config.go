package config

var Version string = "2024.05.08:1.0.3"

var HelpText string = `
Что умеет ресайзер от [by_artisan]:
		
Перебирает все файлы .jpg в папке. В том числе во всех вложенных подпапках (по умолчанию в папке с программой)
Пропорционально вписываеть изображение в два определённых размера (по умалчанию 1200x900 и 186x140)
Если изображение меньше заданного размера, его размер не меняется
Сохраняет полученное изображение добавляя к исходному имени файла символы префикс (по умалчанию b_ и s_)
Исходный файл удаляется

Также программа понимает следующие команды:
--help -h	Вывести подсказку (данное сообщение)
--version -v	Вывести версию программы
--param -p	Изменить параметры по умолчанию
	
Если --param(-p) указан в командах, можно изменить свойства по умолчанию
Prefix_b (pb)		Префикс большого изображения	(b_)
Prefix_s (ps)		Префикс маленького изображения	(s_)
MaxWidthBigImg (mwb)	Ширина большого изображения	(1200)
MaxHightBigImg (mhb)	Высота большого изображения	(900)
MaxWidthSmallImg (mws)	Ширина маленького изображения	(186)
MaxHightSmallImg (mhs)	Высота маленького изображения	(140)
PathRoot (pr)		Корень папки с изображениями	('.')

Формат записи -p pb:prefix_ pr:folder/subfolder
`
