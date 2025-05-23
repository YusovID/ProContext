Условие:

На бесконечной в обе стороны белой полоске размеченной в клеточку находятся два робота. Ровно одна из клеток на полоске - чёрная, и она находится между роботами. Вам необходимо одинаково запрограммировать обоих роботов так, чтоб они встретились. Программа состоит из нескольких строк, каждая из которых содержит ровно одну команду. Допустимые команды: 
1) ML - сделать шаг на клетку влево и перейти к следующей строке программы; 
2) MR - сделать шаг на клетку вправо и перейти к следующей строке программы; 
3) IF FLAG - проверить, находимся ли мы на чёрной клетке. Если да, перейти к следующей строке программы, иначе, перейти к послеследующей строке программы; 
4) GOTO N - перейти к N-й строке программы.
На выполнение каждой из команд, кроме GOTO у робота уходит 1 секунда. GOTO выполняется мгновенно.

Решение:

1) ML;
2) IF_FLAG;
3) GOTO 5;
4) GOTO 1;
5) ML;
6) GOTO 5.

Первый цикл = 1-2-4-1-2-4-... (2 секунды); второй цикл = 1-2-5-6-5-6-... (1 секунда)

Вне зависимости от того, с какой стороны робот будет находиться от черной клетки, он будет двигаться влево со скоростью 1 клетка/2 секунды. Робот, который находился справа от клетки, рано или поздно дойдет до нее и ускорится в 2 раза, т.к. перейдет во второй цикл, в котором нет условия IF_FLAG, которое занимает 1 секунду. Таким образом, он рано или поздно догонит второго робота, который все также двигается со скоростью 1 клетка/2 секунды.

Для визуализации алгоритма, необходимо перейти в ./task_1 и в терминале прописать 

```make run```

или

```go run cmd/main.go```

Для вывода всех доступных команд необходимо прописать 

```make help```