package main

var Base = []string{
	"Deuteronomy 7:9",
	"Deuteronomy 10:12-13",
	"John 1:12-13 ",
	"Romans 11:33-36",
	"Romans 12:1-2",
	"Psalm 56:3-4",
	"Psalm 62:5-7 [8]",
	"Romans 8:1",
	"1 John 2:15-17",
	"Philippians 2:5-7",
	"Philippians 2:8-9",
	"Philippians 2:10-11",
	"Philippians 2:12-13",
	"James 1:2-3",
	"James 1:4-5",
	"Psalm 1:1-2",
	"Psalm 1:3-4",
	"Psalm 1:5-6",
	"Colossians 3:1-3",
	"Ephesians 4:26",
	"Isaiah 40:28-29",
	"Isaiah 40:30-31",
	"Psalm 86:5-7",
	"1 Timothy 2:5",
	"1 Peter 1:3-5 ",
	"Ephesians 6:10-11",
	"Ephesians 6:12-13",
	"Ephesians 6:14-15",
	"Ephesians 6:16-17 [18]",
	"Philippians 1:6",
	"Matthew 10:28",
	"Romans 1:16 [17]",
	"Matthew 11:28-30",
	"Psalm 20:6-7 [8]",
	"James 1:12",
	"2 Corinthians 9:6-7",
	"2 Corinthians 9:8",
	"2 Corinthians 12:9 [10]",
	"Isaiah 64:4",
	"Titus 3:4-6",
	"Isaiah 46:9-10 [11]",
	"Proverbs 1:10",
	"Proverbs 3:5-6 [7]",
	"Proverbs 19:11",
	"John 15:5",
	"John 14:2-3",
	"Psalm 125:1-2",
	"Psalm 141:3-4",
	"1 John 1:8-9",
	"Psalm 23:1-2",
	"Psalm 23:3-4",
	"Psalm 23:5-6",
	"Isaiah 40:8",
	"Romans 10:13-14 [15]",
	"Psalm 16:11",
	"Romans 15:1-2",
	"Psalm 103:1-4",
	"Psalm 103:5-7",
	"Psalm 103:8-10",
	"Psalm 103:11-14",
	"Psalm 103:15-16",
	"Psalm 103:17-19",
	"Psalm 103:20-22",
	"Psalm 86:11",
	"Ephesians 4:29",
	"Ephesians 4:31-32",
	"Deuteronomy 6:4-5",
	"Deuteronomy 6:6-7",
	"2 Corinthians 5:17",
	"Luke 12:32-34 ",
	"Galatians 5:22-23",
	"Galatians 5:24-25",
	"Proverbs 6:20-21",
	"Proverbs 6:22-23",
	"Philippians 4:11-13 ",
	"2 Timothy 1:7",
	"1 Peter 5:6-8",
	"1 Peter 5:9-10 [11]",
	"Proverbs 18:10",
	"Psalm 91:1-2",
	"Psalm 91:3-4",
	"Psalm 91:5-6",
	"Psalm 91:7-8",
	"Psalm 91:9-10",
	"Psalm 91:11-13",
	"Psalm 91:14-16",
	"1 Peter 4:16",
	"John 3:16-17",
	"Acts 4:11-12",
	"Proverbs 29:1,11",
	"Philippians 4:19",
	"1 Corinthians 10:13",
	"Isaiah 53:4-5",
	"Isaiah 53:6",
	"1 Peter 2:24",
	"2 Corinthians 4:17-18",
	"Galatians 2:20 ",
	"Romans 3:23-24",
	"Hebrews 11:6",
	"Romans 14:7-8 [9]",
	"John 3:36",
	"1 Timothy 4:12",
	"1 Corinthians 2:1-2",
	"Revelation 5:12-13",
	"Joshua 1:9",
	"2 Chronicles 16:9",
	"Philippians 3:7-8",
	"Philippians 3:9",
	"Philippians 3:10-11",
	"1 Corinthians 1:18",
	"Psalm 37:[1-2] 3-4",
	"Psalm 37:5-6 [7-8]",
	"Psalm 37:23-24",
	"Hebrews 12:1",
	"Hebrews 12:2 ",
	"Psalm 96:1-3",
	"Psalm 96:4-5",
	"Psalm 96:6-8",
	"Psalm 96:9-10",
	"Isaiah 43:25",
	"Romans 12:9-10",
	"Romans 12:11-13",
	"Romans 12:14-16",
	"Romans 12:17-19",
	"Romans 12:20-21",
	"Proverbs 15:1",
	"James 4:13-14",
	"James 4:15-17",
	"Luke 19:10 ",
	"Psalm 18:30-31",
	"Philippians 4:6-7",
	"Philippians 4:8",
	"Psalm 42:11",
	"Isaiah 46:3-4",
	"Philippians 1:21",
	"Jeremiah 29:11-14",
	"Proverbs 22:1",
	"Psalm 30:4-5",
	"Acts 20:35",
	"Matthew 5:3-6",
	"Matthew 5:7-10",
	"Matthew 5:11-12",
	"1 Corinthians 13:4-7",
	"Psalm 32:8 [9]",
	"Proverbs 31:30",
	"Matthew 6:19-21",
	"1 Corinthians 10:31",
	"Romans 5:18-19",
	"John 5:39-40",
	"1 Peter 2:9-10 [11]",
	"Romans 10:17",
	"Matthew 20:26-28",
	"2 Corinthians 5:21",
	"1 John 3:1 [2]",
	"Ephesians 3:20-21",
	"Matthew 22:37-39",
	"Isaiah 41:10",
	"Isaiah 43:1-3",
	"John 14:6 ",
	"1 Thessalonians 5:14-17",
	"1 Thessalonians 5:18-22",
	"Proverbs 3:9-10",
	"Proverbs 3:11-12",
	"John 10:27-30",
	"Acts 5:29",
	"Psalm 100:1-3",
	"Psalm 100:4-5 ",
	"Ephesians 2:1-3",
	"Ephesians 2:4-5",
	"Ephesians 2:6-7",
	"Ephesians 2:8-10",
	"Romans 5:8",
	"Romans 5:9-10",
	"Psalm 139:1-3",
	"Psalm 139:4-5",
	"Psalm 139:6-8",
	"Psalm 139:9-10",
	"Psalm 139:11-12",
	"Psalm 139:13-14",
	"Psalm 139:15-16",
	"Psalm 139:17-18",
	"Psalm 139:23-24",
	"Proverbs 17:9,22",
	"Hebrews 1:1-2",
	"Hebrews 1:3-4",
	"Jeremiah 1:12",
	"Psalm 9:9-10",
	"James 1:17 ",
	"Psalm 27:1 [2-3]",
	"Psalm 27:4 [5]",
	"Psalm 27:13-14",
	"James 1:22-24",
	"Revelation 2:10",
	"Proverbs 26:20",
	"Jeremiah 9:23-24",
	"Romans 6:23",
	"1 Corinthians 15:1-3",
	"1 Peter 3:18",
	"Psalm 55:22",
	"Psalm 127:1",
	"James 4:7-8",
	"Psalm 119:14-16",
	"John 15:7",
	"Psalm 118:13-14",
	"Proverbs 16:32",
	"Hebrews 3:12-13",
	"1 Corinthians 15:58",
	"John 11:25-26",
	"Daniel 2:20-21",
	"Matthew 9:13",
	"Psalm 121:1-2",
	"Psalm 121:3-4",
	"Psalm 121:5-6",
	"Psalm 121:7-8",
	"Romans 8:28",
	"Romans 8:29-30",
	"Romans 8:31-32",
	"Romans 8:33-34",
	"Romans 8:35-37",
	"Romans 8:38-39",
	"Galatians 6:14",
	"Numbers 23:19",
	"1 Timothy 1:15",
	"Hebrews 13:5-6",
	"Lamentations 3:21-23",
	"Lamentations 3:24-26",
	"Lamentations 3:31-33",
	"Colossians 3:16-17",
	"Isaiah 26:3-4",
	"Psalm 19:7-8",
	"Psalm 19:9-11",
	"John 6:35",
	"Galatians 6:9-10",
	"Psalm 34:1-3",
	"Psalm 34:4-5",
	"Psalm 34:6-8",
	"Psalm 34:9-11",
	"Psalm 34:12-14",
	"Psalm 34:15-16",
	"Psalm 34:17-18",
	"Psalm 34:19-22",
	"Colossians 4:6",
	"John 8:31-32",
	"John 10:10",
	"Jeremiah 32:40",
	"Psalm 73:25-26",
	"Proverbs 4:23-24",
	"Proverbs 4:25-27",
	"James 1:19-20",
	"2 Corinthians 8:9",
	"Psalm 77:13-14",
	"Psalm 118:5-8",
	"1 Timothy 6:6-7",
	"Psalm 79:9",
	"Psalm 84:10-11 [12]",
	"1 John 4:4",
	"1 Corinthians 15:51-52 ",
	"Revelation 21:3",
	"Revelation 21:4",
	"Revelation 21:5-6 [7]",
}
