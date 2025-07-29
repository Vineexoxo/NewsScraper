import React from 'react';

export default function NewsCard({ title, subtitle, image }) {
  return (
    <div className="flex-none w-[320px] md:w-[400px] rounded-3xl overflow-hidden relative group transform transition-all duration-300 hover:scale-[1.02] shadow-md">
      <img src={image} alt={title} className="h-64 w-full object-cover opacity-80" />
      <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-[#0b0f10] via-[#0b0f10cc] to-transparent p-4">
        <h2 className="text-lg font-semibold text-white">{title}</h2>
        <p className="text-sm text-gray-400 mt-1">{subtitle}</p>
      </div>
    </div>
  );
}
