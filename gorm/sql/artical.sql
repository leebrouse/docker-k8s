CREATE TABLE article (
    id INT PRIMARY KEY AUTO_INCREMENT,  -- 文章ID
    title VARCHAR(255) NOT NULL,        -- 文章标题
    cate_id INT NOT NULL,               -- 分类ID（外键）
    state TINYINT DEFAULT 1,            -- 状态（1: 发布, 0: 草稿）
    FOREIGN KEY (cate_id) REFERENCES article_cate(id) ON DELETE CASCADE
);

INSERT INTO article (id, title, cate_id, state) VALUES
(1, '8月份CPI同比上涨2.8% 猪肉价格上涨46.7%', 1, 1),
(2, '中国联通与中国电信共建共享5G网络 用户日感不强', 1, 1),
(3, '林郑月娥表示极度愤慨和强烈谴责不配合警方安全的暴徒', 2, 1),
(4, '这些老师的口头禅，想起那些年“被支配”的恐惧了', 3, 1),
(5, '盘点登喜路一号兰精理事,钟楚红被赞：令人惊奇', 3, 1),
(6, '这些老师的口头禅，想起那些年“被支配”的恐惧', 4, 1);
