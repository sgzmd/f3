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

	ListEntriesFakeResponse = `entry: {
  key: { entity_type: ENTRY_TYPE_SERIES  entity_id: 79232  user_id: "sgzmd-184448015" }  
  entry_name: "Кодекс Охотника"  num_entries: 7  entry_author: "Винокуров, Юрий "  
  book: { book_name: "Кодекс Охотника. Книга I"  book_id: 692730  order_in_sequence: 1 }  
  book: { book_name: "Кодекс Охотника. Книга II"  book_id: 692731  order_in_sequence: 2 }  
  book: { book_name: "Кодекс Охотника. Книга III"  book_id: 710418  order_in_sequence: 3 }  
  book: { book_name: "Кодекс Охотника. Книга IV"  book_id: 701371  order_in_sequence: 4 }  
  book: { book_name: "Кодекс Охотника. Книга V"  book_id: 703519  order_in_sequence: 5 }  
  book: { book_name: "Кодекс Охотника. Книга VI"  book_id: 708175  order_in_sequence: 6 }  
  book: { book_name: "Кодекс Охотника. Книга VII"  book_id: 711534  order_in_sequence: 7 }  
  saved: { seconds: 1677939420 } }  
entry: { key: { entity_type: ENTRY_TYPE_AUTHOR  entity_id: 109170  user_id: "sgzmd-184448015" }  
  entry_name: "Метельский, Николай Александрович"  
  num_entries: 29  
  entry_author: "Метельский, Николай Александрович"  
  book: { book_name: "Клан у пропасти"  book_id: 482165  order_in_sequence: 1 }  
  book: { book_name: "Некоторые пояснения"  book_id: 332587  order_in_sequence: 1 }  
  book: { book_name: "Навыки и техники"  book_id: 332588  order_in_sequence: 1 }  
  book: { book_name: "Имена"  book_id: 332589  order_in_sequence: 1 }  
  book: { book_name: "Боевая техника производства разных стран" book_id: 332590  order_in_sequence: 1 }  
  book: { book_name: "Чужие маски"  book_id: 452501  order_in_sequence: 1 }  
  book: { book_name: "Теряя маски"  book_id: 452502  order_in_sequence: 1 }  
  book: { book_name: "Меняя маски"  book_id: 452503  order_in_sequence: 1 }  
  book: { book_name: "Клан у пропасти (СИ)"  book_id: 471218  order_in_sequence: 1 }  
  book: { book_name: "Юнлинг"  book_id: 504679  order_in_sequence: 1 }  
  book: { book_name: "Удерживая маску"  book_id: 513628  order_in_sequence: 1 }  
  book: { book_name: "Срывая маски"  book_id: 517316  order_in_sequence: 1 }  
  book: { book_name: "Призрачный ученик"  book_id: 530269  order_in_sequence: 1 }  
  book: { book_name: "Маска зверя"  book_id: 530624  order_in_sequence: 1 }  
  book: { book_name: "Унесенный ветром"  book_id: 552329  order_in_sequence: 1 }  
  book: { book_name: "Осколки маски"  book_id: 569518  order_in_sequence: 1 }  
  book: { book_name: "Тень маски"  book_id: 603180  order_in_sequence: 1 }  
  book: { book_name: "Устав от масок"  book_id: 605769  order_in_sequence: 1 }  
  book: { book_name: "Без масок"  book_id: 625424  order_in_sequence: 1 }  
  book: { book_name: "Новый враг"  book_id: 644955  order_in_sequence: 1 }  
  book: { book_name: "Охота на маску. Часть вторая"  book_id: 645426  order_in_sequence: 1 }  
  book: { book_name: "Охота на маску. Часть первая"  book_id: 649660  order_in_sequence: 1 }  
  book: { book_name: "Унесенный ветром"  book_id: 657948  order_in_sequence: 1 }  
  book: { book_name: "Убивая маску"  book_id: 676782  order_in_sequence: 1 }  
  book: { book_name: "Охота на маску. Разбивая иллюзии"  book_id: 692381  order_in_sequence: 1 }  
  book: { book_name: "Убивая маску"  book_id: 704270  order_in_sequence: 1 }  
  book: { book_name: "Убивая маску"  book_id: 710622  order_in_sequence: 1 }  
  book: { book_name: "Унесенный ветром"  book_id: 711519  order_in_sequence: 1 }  
  book: { book_name: "Убивая маску"  book_id: 711966  order_in_sequence: 1 }  
  saved: { seconds: 1677692371 } 
}  

`
)
