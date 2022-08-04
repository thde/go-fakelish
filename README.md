# unamegen

User name generator

## Install

Executable binaries are available in [GitHub Releases](https://github.com/thde/unamegen/releases)

## Usage

```bash
$ unamegen --help
Usage of unamegen [ amount ]:
  -c, --capitalize     Start the usernames with a capital letter
  -h, --help           print help message
  -i, --input string   Input file
      --max int        max length of fake word (default 12)
      --min int        min length of fake word (default 6)
  -1, --no-columns     Don't print the generated usernames in columns
```

## Examples

```bash
$ unamegen
ptical
senturical
calizz
atical
tercon
aricaing
cowful
mineving
dratores
counimper
```

## Supply you own word lists

`unamegen` integrates a basic english word list it uses to generate the random words.

You can use your own word list by using the `-i` flag. This allows you to generate words, that sound like other languages.

<table>
<tr>
  <th>Language</th>
  <th>Example</th>
</tr>
<tr>
  <td>German</td>
  <td>

```shell
curl -s https://web.archive.org/web/20090909075401id_/http://wortschatz.uni-leipzig.de/Papers/top10000de.txt | iconv -f ISO-8859-1 -t UTF-8//TRANSLIT | grep -v '-'
```

  </td>
</tr>
<tr>
  <td>French</td>
  <td>

```shell
curl -s https://web.archive.org/web/20090904105851id_/http://wortschatz.uni-leipzig.de/Papers/top10000fr.txt | iconv -f ISO-8859-1 -t UTF-8//TRANSLIT | grep -v \'
```

  </td>
</tr>
<tr>
  <td>Dutch</td>
  <td>

```shell
curl -s https://web.archive.org/web/20090904014314id_/http://wortschatz.uni-leipzig.de/Papers/top10000nl.txt | iconv -f ISO-8859-1 -t UTF-8//TRANSLIT | grep -v '-'
```

  </td>
</tr>
</table>
