import { Card } from "@/components/retroui/Card";
import { Text } from "@/components/retroui/Text";
import { useLoaderData } from "react-router-dom";

export function NewsPage() {
    let data=useLoaderData();
    console.log("data", data);
    let article=data[0];
    console.log("article", article)
    const raw=article?.description + article?.content;
    const sections = raw.split("/n").map(s => s.trim()).filter(Boolean);
    return (
        
        <Card>
            <Card.Header>
                <Card.Title>{article.title}</Card.Title>
                    <Card.Description>{article.pubDate}</Card.Description>
            </Card.Header>
            <Card.Content>
                    <div className="space-y-4 p-6">
                        {sections.map((section, i) => (
                            <div key={i} >
                                <article className="mb-6">
                                    <Text className="text-lg leading-relaxed">{section}</Text>
                                </article>
                            </div>
                        ))}
                    </div>    
            </Card.Content>
            <footer className="p-6">
                    {article.keywords?.length > 0 && (
                        <div className="mb-6">
                            <h2 className="text-sm font-semibold text-gray-600 mb-2">Keywords:</h2>
                            <div className="flex flex-wrap gap-2">
                                {article.keywords.map((kw, i) => (
                                    <span
                                        key={i}
                                        className="px-3 py-1 bg-gray-100 rounded-full text-sm text-gray-700"
                                    >
                                        {kw}
                                    </span>
                                ))}
                            </div>
                        </div>
                    )}

                    {/* External Link */}
                    <footer>
                        <a
                            href={article.link}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="inline-block px-4 py-2 bg-blue-600 text-white rounded-lg 
                    shadow hover:bg-blue-700 transition-colors"
                        >
                            Read Full Article →
                        </a>
                    </footer>
            </footer>
            </Card>
        
    );
}