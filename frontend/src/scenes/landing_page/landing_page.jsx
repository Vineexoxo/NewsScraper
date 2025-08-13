/* eslint-disable no-unused-vars */
import React from "react";
import { motion } from "framer-motion";
import { ArticleCard } from "../../components/catalog_item/article_card";

// Single-file React component (TailwindCSS required in the project)
// Usage: paste into a Vite/CRA React project and ensure Tailwind + Framer Motion are installed.

const ARTICLES = [
  {
    id: 1,
    title: "The Return of Quiet Cities: How Urban Life is Rebalancing",
    byline: "By A. Reporter",
    excerpt:
      "A sweeping look at the ways cities are reshaping post-pandemic life — transport, markets and community-led design.",
    time: "Aug 12, 2025",
  },
  {
    id: 2,
    title: "Paper & Pixel: The New Rules of Typography",
    byline: "By M. Type",
    excerpt:
      "Designers are borrowing from the old press to create new digital reading experiences that feel comforting and tactile.",
    time: "Aug 9, 2025",
  },
  {
    id: 3,
    title: "Markets in Sepia: Why Monochrome Still Sells",
    byline: "By E. Finance",
    excerpt:
      "A look into how reduced color palettes affect brand trust, attention, and long-form reading online.",
    time: "Aug 1, 2025",
  },
];


export default function NewspaperLikeSite() {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-1 gap-4">
      {ARTICLES.map(({id,title,byline,excerpt,time}) => (
        <ArticleCard key={id} text={excerpt} />
      ))}


    </div>
  );
}
