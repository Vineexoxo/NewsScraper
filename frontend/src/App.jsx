import React from 'react';
import NewsCard from './Components/NewsCard/NewsCard.jsx';

const dummyNews = [
  {
    id: 1,
    title: "Neural Clouds Take Over Memory Backups",
    subtitle: "A new era of encrypted recall through organic storage.",
    image: "https://source.unsplash.com/1600x900/?neural,technology",
  },
  {
    id: 2,
    title: "Orbital Internet Mesh Launched",
    subtitle: "Satellite-powered broadband goes full decentralized.",
    image: "https://source.unsplash.com/1600x900/?space,internet",
  },
];

export default function App() {
  return (
    <div className="bg-[#0b0f10] min-h-screen text-white font-light tracking-wide">
      <header className="text-center pt-6 text-gray-400 text-sm uppercase tracking-[0.2em]">
        NEO / FEED
      </header>
      <section className="flex overflow-x-auto gap-6 p-8 hide-scrollbar">
        {dummyNews.map(article => (
          <NewsCard key={article.id} {...article} />
        ))}
      </section>
    </div>
  );
}
