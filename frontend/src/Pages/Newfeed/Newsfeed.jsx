import React, { useState, useEffect } from 'react';
import { Clock, ArrowUpRight, Zap, Eye, Brain } from 'lucide-react';

const NewsFeed = () => {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [activeIndex, setActiveIndex] = useState(0);

  const mockArticles = [
    {
      id: 1,
      title: "Neural Interfaces Merge Human Consciousness with AI",
      excerpt: "Breakthrough technology enables direct brain-computer communication, revolutionizing how we interact with digital systems.",
      author: "Dr. Aria Nexus",
      publishedAt: "2h",
      readTime: "7 min",
      image: "https://images.unsplash.com/photo-1677442136019-21780ecad995?w=1200&h=600&fit=crop",
      priority: "CRITICAL",
      engagement: 2847
    },
    {
      id: 2,
      title: "Quantum Processors Achieve 99.9% Error Correction",
      excerpt: "Major leap in quantum computing stability brings us closer to solving humanity's greatest computational challenges.",
      author: "Zara Chen-47",
      publishedAt: "4h",
      readTime: "5 min",
      image: "https://images.unsplash.com/photo-1635070041078-e363dbe005cb?w=1200&h=600&fit=crop",
      priority: "HIGH",
      engagement: 1923
    },
    {
      id: 3,
      title: "Fusion Reactors Power First Carbon-Negative City",
      excerpt: "Breakthrough in clean fusion energy enables the world's first city to achieve negative carbon emissions.",
      author: "Marcus Sol",
      publishedAt: "6h",
      readTime: "9 min",
      image: "https://images.unsplash.com/photo-1509391366360-2e959784a276?w=1200&h=600&fit=crop",
      priority: "HIGH",
      engagement: 3156
    },
    {
      id: 4,
      title: "Holographic Displays Replace All Screens by 2030",
      excerpt: "Spatial computing advances make traditional displays obsolete as holographic interfaces become mainstream.",
      author: "Eva Matrix",
      publishedAt: "8h",
      readTime: "6 min",
      image: "https://images.unsplash.com/photo-1522202176988-66273c2fd55f?w=1200&h=600&fit=crop",
      priority: "MEDIUM",
      engagement: 1674
    },
    {
      id: 5,
      title: "Synthetic Biology Creates Self-Repairing Materials",
      excerpt: "Living materials that heal themselves revolutionize construction and manufacturing industries worldwide.",
      author: "Dr. Bio-X",
      publishedAt: "12h",
      readTime: "8 min",
      image: "https://images.unsplash.com/photo-1416879595882-3373a0480b5b?w=1200&h=600&fit=crop",
      priority: "MEDIUM",
      engagement: 987
    }
  ];

  useEffect(() => {
    setTimeout(() => {
      setArticles(mockArticles);
      setLoading(false);
    }, 2000);
  }, []);

  useEffect(() => {
    const interval = setInterval(() => {
      setActiveIndex(prev => (prev + 1) % articles.length);
    }, 4000);
    return () => clearInterval(interval);
  }, [articles.length]);

  if (loading) {
    return (
      <div className="min-h-screen bg-black flex items-center justify-center relative overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-br from-blue-900/20 via-purple-900/20 to-pink-900/20"></div>
        <div className="absolute inset-0" style={{
          backgroundImage: `radial-gradient(circle at 20% 50%, rgba(120, 119, 198, 0.3) 0%, transparent 50%),
                           radial-gradient(circle at 80% 20%, rgba(255, 119, 198, 0.3) 0%, transparent 50%),
                           radial-gradient(circle at 40% 80%, rgba(120, 219, 255, 0.3) 0%, transparent 50%)`
        }}></div>
        
        <div className="relative z-10 text-center space-y-6">
          <div className="w-20 h-20 mx-auto relative">
            <div className="absolute inset-0 rounded-full border-2 border-cyan-400/30"></div>
            <div className="absolute inset-2 rounded-full border-2 border-purple-400/30 animate-spin"></div>
            <div className="absolute inset-4 rounded-full border-2 border-pink-400/30 animate-pulse"></div>
            <div className="absolute inset-6 rounded-full bg-gradient-to-r from-cyan-400 to-purple-400 animate-pulse"></div>
          </div>
          <div className="space-y-2">
            <div className="text-2xl font-light text-cyan-300 tracking-wider">NEURAL FEED</div>
            <div className="text-sm text-gray-400 font-mono">Initializing quantum data stream...</div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-black relative overflow-hidden">
      {/* Animated Background */}
      <div className="fixed inset-0 z-0">
        <div className="absolute inset-0 bg-gradient-to-br from-blue-900/10 via-purple-900/10 to-pink-900/10"></div>
        <div className="absolute inset-0" style={{
          backgroundImage: `radial-gradient(circle at 20% 50%, rgba(120, 119, 198, 0.15) 0%, transparent 50%),
                           radial-gradient(circle at 80% 20%, rgba(255, 119, 198, 0.15) 0%, transparent 50%),
                           radial-gradient(circle at 40% 80%, rgba(120, 219, 255, 0.15) 0%, transparent 50%)`
        }}></div>
        <div className="absolute inset-0 bg-[url('data:image/svg+xml,%3Csvg opacity-20" 
        width="60" height="60" viewBox="0 0 60 60" xmlns="http://www.w3.org/2000/svg"
        fill-rule="evenodd" 
        fill="%23ffffff" fill-opacity="0.02"  path d="M30 30c0-16.569-13.431-30-30-30v60c16.569 0 30-13.431 30-30z" />
        </div>

      {/* Header */}
      <header className="relative z-40 border-b border-cyan-500/20 backdrop-blur-xl bg-black/50">
        <div className="max-w-7xl mx-auto px-6 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <div className="relative">
                <div className="w-10 h-10 rounded-lg bg-gradient-to-br from-cyan-400 via-blue-500 to-purple-600 flex items-center justify-center">
                  <Brain className="w-5 h-5 text-white" />
                </div>
                <div className="absolute -top-1 -right-1 w-3 h-3 bg-green-400 rounded-full animate-pulse"></div>
              </div>
              <div>
                <h1 className="text-xl font-light text-cyan-300 tracking-wider">NEURAL FEED</h1>
                <div className="text-xs text-gray-500 font-mono">v4.2.1 // QUANTUM PROTOCOL</div>
              </div>
            </div>
            
            <div className="flex items-center space-x-6">
              <div className="hidden md:flex items-center space-x-4 text-xs font-mono">
                <span className="text-green-400">◉ LIVE</span>
                <span className="text-gray-400">|</span>
                <span className="text-cyan-400">{articles.length} FEEDS</span>
                <span className="text-gray-400">|</span>
                <span className="text-purple-400">12.4K NODES</span>
              </div>
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="relative z-10 max-w-7xl mx-auto px-6 py-8">
        <div className="space-y-8">
          {articles.map((article, index) => (
            <ArticleCard
              key={article.id}
              article={article}
              index={index}
              isActive={index === activeIndex}
            />
          ))}
        </div>
      </main>

      {/* Floating Navigation */}
      <div className="fixed right-6 top-1/2 transform -translate-y-1/2 z-50 space-y-2">
        {articles.map((_, index) => (
          <button
            key={index}
            onClick={() => setActiveIndex(index)}
            className={`w-2 h-8 rounded-full transition-all duration-300 ${
              index === activeIndex
                ? 'bg-gradient-to-b from-cyan-400 to-purple-400'
                : 'bg-gray-700 hover:bg-gray-600'
            }`}
          />
        ))}
      </div>
    </div>
  );
};

const ArticleCard = ({ article, index, isActive }) => {
  const getPriorityColor = (priority) => {
    switch(priority) {
      case 'CRITICAL': return 'from-red-500 to-orange-500';
      case 'HIGH': return 'from-yellow-500 to-orange-500';
      default: return 'from-blue-500 to-cyan-500';
    }
  };

  return (
    <article
      className={`group relative animate-fade-in transform transition-all duration-1000 ${
        isActive ? 'scale-105' : 'scale-100'
      }`}
      style={{ animationDelay: `${index * 200}ms` }}
    >
      <div className="relative overflow-hidden rounded-2xl border border-gray-800/50 bg-gray-900/20 backdrop-blur-xl hover:border-cyan-500/30 transition-all duration-500">
        {/* Priority Indicator */}
        <div className="absolute top-4 left-4 z-20">
          <div className={`px-3 py-1 rounded-full bg-gradient-to-r ${getPriorityColor(article.priority)} text-white text-xs font-mono font-bold tracking-wider`}>
            {article.priority}
          </div>
        </div>

        {/* Engagement Metrics */}
        <div className="absolute top-4 right-4 z-20 flex items-center space-x-2">
          <div className="flex items-center space-x-1 px-2 py-1 rounded-full bg-black/50 backdrop-blur-sm">
            <Eye className="w-3 h-3 text-cyan-400" />
            <span className="text-xs text-cyan-400 font-mono">{article.engagement}</span>
          </div>
        </div>

        <div className="lg:flex">
          <div className="lg:w-1/2 relative">
            <div className="relative overflow-hidden h-80 lg:h-96">
              <img
                src={article.image}
                alt={article.title}
                className="w-full h-full object-cover group-hover:scale-110 transition-transform duration-700"
              />
              <div className="absolute inset-0 bg-gradient-to-t from-black via-black/20 to-transparent"></div>
              
              {/* Holographic Effect */}
              <div className="absolute inset-0 opacity-30">
                <div className="absolute inset-0 bg-gradient-to-r from-transparent via-cyan-500/10 to-transparent transform skew-x-12 animate-pulse"></div>
              </div>
            </div>
          </div>
          
          <div className="lg:w-1/2 p-8 space-y-6">
            <div className="flex items-center space-x-4 text-xs font-mono">
              <div className="flex items-center space-x-2">
                <div className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></div>
                <span className="text-cyan-400">@{article.author}</span>
              </div>
              <span className="text-gray-500">|</span>
              <div className="flex items-center space-x-1 text-gray-400">
                <Clock className="w-3 h-3" />
                <span>{article.publishedAt}</span>
              </div>
              <span className="text-gray-500">|</span>
              <div className="flex items-center space-x-1 text-purple-400">
                <Zap className="w-3 h-3" />
                <span>{article.readTime}</span>
              </div>
            </div>
            
            <div className="space-y-4">
              <h2 className="text-2xl lg:text-3xl font-light text-white leading-tight group-hover:text-cyan-300 transition-colors duration-300">
                {article.title}
              </h2>
              <p className="text-gray-300 leading-relaxed text-lg">
                {article.excerpt}
              </p>
            </div>

            <div className="flex items-center justify-between pt-6 border-t border-gray-800/50">
              <div className="flex items-center space-x-4">
                <button className="px-4 py-2 rounded-lg bg-gradient-to-r from-cyan-500/20 to-blue-500/20 border border-cyan-500/30 text-cyan-300 hover:from-cyan-500/30 hover:to-blue-500/30 transition-all duration-300 text-sm font-mono">
                  ANALYZE
                </button>
                <button className="px-4 py-2 rounded-lg bg-gradient-to-r from-purple-500/20 to-pink-500/20 border border-purple-500/30 text-purple-300 hover:from-purple-500/30 hover:to-pink-500/30 transition-all duration-300 text-sm font-mono">
                  ENHANCE
                </button>
              </div>
              
              <button className="group/read flex items-center space-x-2 px-6 py-3 rounded-lg bg-gradient-to-r from-cyan-500 to-purple-500 text-white hover:from-cyan-400 hover:to-purple-400 transition-all duration-300 font-mono text-sm">
                <span>INTERFACE</span>
                <ArrowUpRight className="w-4 h-4 group-hover/read:translate-x-1 group-hover/read:-translate-y-1 transition-transform duration-300" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </article>
  );
};

export default NewsFeed;