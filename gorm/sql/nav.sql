
/* type Navs struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Url     string    `json:"url"`
	Status  int       `json:"status"`
	Sort    int       `json:"sort"`
	AddTime time.Time `json:"add_time"`
} */
CREATE Table IF NOT EXISTS navs(
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255),
    url VARCHAR(255),
    status INT,
    sort INT,
    add_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

/* insert values */
INSERT INTO navs (title, url, status, sort) VALUES
('首页', 'https://example.com/home', 1, 1),
('关于我们', 'https://example.com/about', 1, 2),
('服务', 'https://example.com/services', 1, 3),
('产品', 'https://example.com/products', 1, 4),
('博客', 'https://example.com/blog', 1, 5),
('联系我们', 'https://example.com/contact', 1, 6),
('隐私政策', 'https://example.com/privacy', 0, 7),
('帮助中心', 'https://example.com/help', 1, 8);
