package rpc

const (
	GlobalSearchFakeResponse = `
original_request: {
  search_term: "Маски"
}
entry: {
  entry_type: ENTRY_TYPE_AUTHOR
  entry_name: "Роман Маскин"
  author: "Роман Маскин"
  entry_id: 7944
  num_entities: 1
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Без маски"
  author: "Генри Лайон Олди"
  entry_id: 12644
  num_entities: 2
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Брэдбери, Рэй. Сборники рассказов: 21. Маски"
  author: "Рэй Брэдбери"
  entry_id: 22777
  num_entities: 7
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Вечные маски"
  author: "Анна Розен"
  entry_id: 74191
  num_entities: 2
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Врачи без маски: реальные истории"
  author: "Алексей Александрович Виленский, Арсений Кожухов, Дмитрий Николаевич Кушкин, Ярослав Андреевич Соколов"
  entry_id: 47627
  num_entities: 4
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Маски (Медведева)"
  author: "Алена Викторовна Медведева"
  entry_id: 48975
  num_entities: 2
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Маски [= Унесенный ветром]"
  author: "Николай Александрович Метельский"
  entry_id: 34145
  num_entities: 17
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Маски [Брайт]"
  author: "Фреда Брайт"
  entry_id: 52698
  num_entities: 3
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Маски [Лана Черная]"
  author: "Лана Черная"
  entry_id: 64911
  num_entities: 2
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Маски врагов"
  author: "Владислав Валерьевич Выставной"
  entry_id: 27837
  num_entities: 2
}
entry: {
  entry_type: ENTRY_TYPE_SERIES
  entry_name: "Маски смерти"
  author: "Семен Глинский"
  entry_id: 32206
  num_entities: 1
}
`
	TrackEntryFakeRequer = `
key:  {
  entity_type:  ENTRY_TYPE_SERIES
  entity_id:  34145
  user_id:  "user"
}`

	ListEntriesFakeResponse = `
entry:  {
  entry_type:  ENTRY_TYPE_SERIES
  entry_name:  "Маски [= Унесенный ветром]"
  entry_id:  34145
  num_entries:  17
  user_id:  "user"
  book:  {
    book_name:  "Маска зверя"
    book_id:  530624
  }
  book:  {
    book_name:  "Удерживая маску"
    book_id:  513628
  }
  book:  {
    book_name:  "Срывая маски"
    book_id:  517316
  }
  book:  {
    book_name:  "Теряя маски"
    book_id:  452502
  }
  book:  {
    book_name:  "Чужие маски"
    book_id:  452501
  }
  book:  {
    book_name:  "Меняя маски"
    book_id:  452503
  }
  book:  {
    book_name:  "Унесенный ветром"
    book_id:  552329
  }
  book:  {
    book_name:  "Осколки маски"
    book_id:  569518
  }
  book:  {
    book_name:  "Устав от масок"
    book_id:  605769
  }
  book:  {
    book_name:  "Тень маски"
    book_id:  603180
  }
  book:  {
    book_name:  "Без масок"
    book_id:  625424
  }
  book:  {
    book_name:  "Новый враг"
    book_id:  644955
  }
  book:  {
    book_name:  "Унесенный ветром"
    book_id:  657948
  }
  book:  {
    book_name:  "Охота на маску. Часть вторая"
    book_id:  645426
  }
  book:  {
    book_name:  "Охота на маску. Часть первая"
    book_id:  649660
  }
}`
)
