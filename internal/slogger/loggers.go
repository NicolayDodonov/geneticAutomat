package slogger

import "log/slog"

// Логгер срока жизни мира
var LogWorldAge *slog.Logger

var LogWorldAgeCSV *slog.Logger

// Логгер получения инфо о endPopulation's ботах и их геноме с перезаписью
var LogWorldBest *slog.Logger

// Логгер получения инфо о поведении ботов
var LogEntityInfo *slog.Logger

// Логгер ошибок
var LogErrors *slog.Logger
