INSERT INTO articles (
    article_id, title, link, keywords, creator, description, content, pub_date, pub_date_tz, 
    image_url, video_url, source_id, source_name, source_priority, source_url, source_icon, 
    language, country, category, sentiment, sentiment_stats, ai_tag, ai_region, ai_org, 
    ai_summary, ai_content, duplicate
) VALUES
(
    'a1', 'AI Revolution in 2025', 'https://news.com/ai-2025', 
    ARRAY['AI','Technology','Future'], ARRAY['John Doe','Jane Smith'], 
    'A deep dive into AI advancements in 2025', 'Full article content here...', 
    '2025-08-10', 'UTC+5:30', 'https://images.com/ai2025.jpg', NULL, 
    'src1', 'Tech News', 1, 'https://technews.com', 'https://icons.com/technews.png', 
    'en', ARRAY['US','IN'], ARRAY['Technology','Science'], 'positive', '80% positive, 15% neutral, 5% negative', 
    'AI Trends', 'Asia', 'OpenAI', 'Summary of AI advancements...', 
    'Processed AI content here...', FALSE
),
(
    'a2', 'SpaceX Mars Mission Update', 'https://news.com/spacex-mars', 
    ARRAY['SpaceX','Mars','Space'], ARRAY['Elon Musk'], 
    'Latest updates on Mars colonization plans', 'Detailed article content here...', 
    '2025-08-09', 'UTC+0', NULL, NULL, 
    'src2', 'Space News', 2, 'https://spacenews.com', 'https://icons.com/spacenews.png', 
    'en', ARRAY['US'], ARRAY['Space','Exploration'], 'neutral', '50% positive, 40% neutral, 10% negative', 
    'Space Exploration', 'Mars', 'SpaceX', 'Brief summary...', 
    'Full processed content...', FALSE
),
(
    'a3', 'Economic Outlook 2026', 'https://news.com/economy-2026', 
    ARRAY['Economy','Finance'], NULL, 
    'Predictions for the global economy in 2026', 'Economic trends and forecasts...', 
    '2025-08-08', 'UTC-5', NULL, NULL, 
    'src3', 'Finance Daily', 3, 'https://financedaily.com', 'https://icons.com/finance.png', 
    'en', ARRAY['US','UK'], ARRAY['Economy','Finance'], 'negative', '30% positive, 20% neutral, 50% negative', 
    'Economic Trends', 'Global', 'IMF', 'Short summary...', 
    'Processed content here...', TRUE
);