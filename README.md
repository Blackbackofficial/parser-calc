# parser-async
# Домашнее задание 1

## Дисклеймер

Это задание состоит из 2х частей, которые нужно сдавать вместе.
Отдельно первая часть, как и отдельно вторая часть, не оценивается
как половина задания.

Задачи включают в себя как написание функциональности, так и её
тестирование.

Все домашние задания должны выполняться в приватных репозиториях.

## Часть 1. Uniq

Нужно реализовать утилиту, с помощью которой можно вывести или отфильтровать
повторяющиеся строки в файле (аналог UNIX утилиты `uniq`). Причём повторяющиеся
входные строки не должны распозноваться, если они не следуют строго друг за другом.
Сама утилита имеет набор параметров, которые необходимо поддержать.

### Параметры

`-с` - подсчитать количество встречаний строки во входных данных.
Вывести это число перед строкой отделив пробелом.

`-d` - вывести только те строки, которые повторились во входных данных.

`-u` - вывести только те строки, которые не повторились во входных данных.

`-f num_fields` - не учитывать первые `num_fields` полей в строке.
Полем в строке является непустой набор символов отделённый пробелом.

`-s num_chars` - не учитывать первые `num_chars` символов в строке.
При использовании вместе с параметром `-f` учитываются первые символы
после `num_fields` полей (не учитывая пробел-разделитель после
последнего поля).

`-i` - не учитывать регистр букв.

### Использование

`uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]`

1. Все параметры опциональны. Поведения утилиты без параметров --
   простой вывод уникальных строк из входных данных.

2. Параметры c, d, u взаимозаменяемы. Необходимо учитывать,
   что параллельно эти параметры не имеют никакого смысла. При
   передаче одного вместе с другим нужно отобразить пользователю
   правильное использование утилиты

3. Если не передан input_file, то входным потоком считать stdin

4. Если не передан output_file, то выходным потоком считать stdout

### Пример работы

<details>
    <summary>Без параметров</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
$cat input.txt | go run uniq.go
I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.
```

</details>

<details>
    <summary>С параметром input_file</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
$go run uniq.go input.txt
I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.
```

</details>

<details>
    <summary>С параметрами input_file и output_file</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
$go run uniq.go input.txt output.txt
$cat output.txt
I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.
```

</details>

<details>
    <summary>С параметром -c</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
$cat input.txt | go run uniq.go -c
3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
2 I love music of Kartik.
```

</details>

<details>
    <summary>С параметром -d</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
$cat input.txt | go run uniq.go -d
I love music.
I love music of Kartik.
I love music of Kartik.
```

</details>

<details>
    <summary>С параметром -u</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
$cat input.txt | go run uniq.go -u

Thanks.
```

</details>

<details>
    <summary>С параметром -i</summary>

```bash
$cat input.txt
I LOVE MUSIC.
I love music.
I LoVe MuSiC.

I love MuSIC of Kartik.
I love music of kartik.
Thanks.
I love music of kartik.
I love MuSIC of Kartik.
$cat input.txt | go run uniq.go -i
I LOVE MUSIC.

I love MuSIC of Kartik.
Thanks.
I love music of kartik.
```

</details>

<details>
    <summary>С параметром -f num</summary>

```bash
$cat input.txt
We love music.
I love music.
They love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
$cat input.txt | go run uniq.go -f 1
We love music.

I love music of Kartik.
Thanks.
```

</details>

<details>
    <summary>С параметром -s num</summary>

```bash
$cat input.txt
I love music.
A love music.
C love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
$cat input.txt | go run uniq.go -s 1
I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
```

</details>

### Тестирование

Нужно протестировать поведение написанной функциональности
с различными параметрами. Для тестирования нужно написать unit-тесты
на эту функциональность. Тесты нужны как для успешных случаев,
так и для неуспешных. Примеры с тестами мы будем показывать ещё на
следующих лекциях, но сейчас можно посмотреть в [шестом примере первой лекции](https://github.com/go-park-mail-ru/lectures/blob/master/1-basics/6_is_sorted/sorted/sorted_test.go).

### Материалы в помощь

В `1-basics/readme.md` есть список книг по го, а так же по всем частым и нужным операциям, там вы можете найти многие примеры кода, которые вам пригодятся.

Материалы в помощь:

* https://habrahabr.ru/post/306914/ - пакет io

* https://golang.org/pkg/sort/

* https://golang.org/pkg/io/

* https://golang.org/pkg/io/ioutil/

* https://godoc.org/flag - пакет для флагов

* https://godoc.org/github.com/stretchr/testify - удобный набор
  пакетов для тестирования

* https://golang.org/pkg/bufio/#Scanner - удобный способ прочитать
  линии из потока данных

### Best practices

1. Уникализация строк может понадобиться не только как утилита,
   но и как часть более крупной логики. Для этого саму функцию
   уникализации можно вынести в отдельный пакет. Поскольку
   более крупная логика не всегда связана с чтением аргументов
   и данных из файла или stdin, то на вход этой функции нужно
   передавать слайс строк и аргументы.

2. Множество параметров, которые вдобавок и опциональны, лучше
   передавать структурой (например Options). Так проще расширять
   функциональность, а внешнему пользователю функции (не всей утилиты)
   будет проще передать правильные аргументы внутрь.

3. Как файл, так и stdin удовлетворяет интерфейсу io.Reader.
   Поэтому логику по чтению можно сделать универсальной. Аналогично
   и с записью -- io.Writer

4. Для написания однотипных тестовых случаев используется
   [табличное тестирование](https://github.com/golang/go/wiki/TableDrivenTests). Получается, что можно написать две функции
   тестов: успешные тестовые случаи и неуспешные тестовые случаи.

5. Для сравнения ожидаемого и действительного можно использовать
   пакет [require](https://godoc.org/github.com/stretchr/testify/require).
   Кроме простых сравнений на равенство пакет предоставляет много
   других ассертов.

6. Тесты не должны зависеть от внешних ресурсов. Не нужно читать
   файлы внутри теста. Так же не нужно тестировать передачу параметров
   при вызове утилиты. Никакого внешнего взаимодействия. Тестирование
   функции должно быть построено на том, что мы передаём некоторые
   входные данные в функцию и сравниваем ответ функции с ожидаемыми
   выходными данными.

## Часть 2. Calc

Нужно написать калькулятор, умеющий вычислять выражение, подаваемое на STDIN.

Достаточно реализовать сложение, вычитание, умножение, деление и поддержку скобок.

Тут также нужны тесты 🙂 Тестами нужно покрыть все операции.

### Пример работы

```bash
    $ go run calc.go "(1+2)-3"
    0

    $ go run calc.go "(1+2)*3"
    9
```


# Домашнее задание №2

------

В этом задании мы пишем аналог unix pipeline, что-то вроде:
```
grep 127.0.0.1 | awk '{print $2}' | sort | uniq -c | sort -nr
```

Когда STDOUT одной программы передаётся как STDIN в другую программу

Но в нашем случае эти роли выполняют каналы, которые мы передаём из одной функции в другую.

*Это сложное задание, не стесняйтесь просить помощи, оно делается не сразу, но когда в голове щёлкнет -  всё становится очень просто*

*Задание при применению материалов лекции. Всё что вам необходимо есть в коде лекции*

Само задание по сути состоит из двух частей
* Написание функции ExecutePipeline которая обеспечивает нам конвейерную обработку функций-воркеров, которые что-то делают.
* Написание нескольких функций, которые считают нам какую-то условную хеш-сумму от входных данных

Расчет хеш-суммы реализован следующей цепочкой:
* SingleHash считает значение crc32(data)+"~"+crc32(md5(data)) ( конкатенация двух строк через ~), где data - то что пришло на вход (по сути - числа из первой функции)
* MultiHash считает значение crc32(th+data)) (конкатенация цифры, приведённой к строке и строки), где th=0..5 ( т.е. 6 хешей на каждое входящее значение ), потом берёт конкатенацию результатов в порядке расчета (0..5), где data - то что пришло на вход (и ушло на выход из SingleHash)
* CombineResults получает все результаты, сортирует (https://golang.org/pkg/sort/), объединяет отсортированный результат через _ (символ подчеркивания) в одну строку
* crc32 считается через функцию DataSignerCrc32
* md5 считается через DataSignerMd5

В чем подвох:
* DataSignerMd5 может одновременно вызываться только 1 раз, считается 10 мс. Если одновременно запустится несколько - будет перегрев на 1 сек
* DataSignerCrc32, считается 1 сек
* На все расчеты у нас 3 сек.
* Если делать в лоб, линейно - для 7 элементов это займёт почти 57 секунд, следовательно надо это как-то распараллелить

Результаты, которые выводятся если отправить 2 значения (закомментировано в тесте):

```
0 SingleHash data 0
0 SingleHash md5(data) cfcd208495d565ef66e7dff9f98764da
0 SingleHash crc32(md5(data)) 502633748
0 SingleHash crc32(data) 4108050209
0 SingleHash result 4108050209~502633748
4108050209~502633748 MultiHash: crc32(th+step1)) 0 2956866606
4108050209~502633748 MultiHash: crc32(th+step1)) 1 803518384
4108050209~502633748 MultiHash: crc32(th+step1)) 2 1425683795
4108050209~502633748 MultiHash: crc32(th+step1)) 3 3407918797
4108050209~502633748 MultiHash: crc32(th+step1)) 4 2730963093
4108050209~502633748 MultiHash: crc32(th+step1)) 5 1025356555
4108050209~502633748 MultiHash result: 29568666068035183841425683795340791879727309630931025356555

1 SingleHash data 1
1 SingleHash md5(data) c4ca4238a0b923820dcc509a6f75849b
1 SingleHash crc32(md5(data)) 709660146
1 SingleHash crc32(data) 2212294583
1 SingleHash result 2212294583~709660146
2212294583~709660146 MultiHash: crc32(th+step1)) 0 495804419
2212294583~709660146 MultiHash: crc32(th+step1)) 1 2186797981
2212294583~709660146 MultiHash: crc32(th+step1)) 2 4182335870
2212294583~709660146 MultiHash: crc32(th+step1)) 3 1720967904
2212294583~709660146 MultiHash: crc32(th+step1)) 4 259286200
2212294583~709660146 MultiHash: crc32(th+step1)) 5 2427381542
2212294583~709660146 MultiHash result: 4958044192186797981418233587017209679042592862002427381542

CombineResults 29568666068035183841425683795340791879727309630931025356555_4958044192186797981418233587017209679042592862002427381542
```

Код писать в signer.go. В этот файл не надо добавлять ничего из common.go, он уже будет на сервере.

Запускать как `go test -v -race`

Подсказки:

* Задание построено так чтобы хорошо разобраться со всем материалом лекции, т.е. вдумчиво посмотреть примеры и применить их на практике. Искать по гуглу или стек оферфлоу ничего не надо
* Вам не надо накапливать данные - сразу передаём их дальше ( например awk из кода выше - на это есть отдельный тест. Разве что функция сама не решает накопить - у нас это CombineResults или sort из кода выше
* Подумайте, как будет организовано завершение функции если данные конечны. Что для этого надо сделать?
* Если вам встретился рейс ( опция -race ) - исследуйте его вывод - когда читаем, когда пишем, из каких строк кода. Там как правило содержится достаточно информации для нахождения источника проблемы.
* Прежде чем приступать к распараллеливанию функций, чтобы уложиться в отведённый таймаут - сначала напишите линейный код, который будет выдавать правильный результат, лучше даже начать с меньшего количества значений чтобы совпадало с тем что в задании
* Вы можете ожидать, что у вас никогда не будет более 100 элементов во входных данных
* Ответ на вопрос "когда закрывается цикл по каналу" помогает в реализации ExecutePipeline
* Ответ на вопрос "мне нужны результаты предыдущих вычислений?" помогают распараллелить SingleHash и MultiHash
* Хорошо помогает нарисовать схему рассчетов
* Естественно нельзя самим считать хеш-суммы в обход предоставляемых функций - их вызов будет проверяться

Эталонное решение занимает 130 строк с учетом дебага который вы видите выше