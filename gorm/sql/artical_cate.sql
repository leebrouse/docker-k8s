CREATE TABLE article_cate (
    id INT PRIMARY KEY AUTO_INCREMENT,  -- 分类ID
    title VARCHAR(255) NOT NULL,        -- 分类名称
    state TINYINT DEFAULT 1             -- 状态（1: 启用, 0: 禁用）
);

INSERT INTO article_cate (id, title, state) VALUES
(1, '新闻', 1),
(2, '科技', 1),
(3, '财经', 1),
(4, '社会', 1);
