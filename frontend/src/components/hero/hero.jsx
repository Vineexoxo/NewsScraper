import { Text } from "@/components/retroui/Text";
import React from "react";

export default function HeroSection() {
    return (
        <>
        
        
        <div className="flex items-center justify-center min-h-screen bg-white relative overflow-hidden">
            {/* Background grid effect */}
            <div className="absolute inset-0 bg-[linear-gradient(to_right,#e5e7eb_1px,transparent_1px),linear-gradient(to_bottom,#e5e7eb_1px,transparent_1px)] bg-[size:40px_40px] opacity-50 pointer-events-none" />

            {/* Content */}
            <div className="relative z-10 text-center px-6">
                <Text as="h1" className="text-5xl md:text-6xl font-extrabold leading-tight">
                    Scrape Websites and make your{" "}
                    <span className="relative inline-block">
                        <span className="relative text-black z-10"> life easy!</span>
                        <span className="absolute inset-0 text-yellow-400 translate-x-1 translate-y-1 -z-0">
                            life easy!
                        </span>
                    </span>
                </Text>

                <p className="mt-6 text-gray-600 max-w-xl mx-auto text-lg">
                    We do the heavy lifting for you, so you can focus on what matters most.
                </p>
            </div>
        </div>
            <section className="bg-pink-200 py-20">
                <div className="text-center mb-12">
                    <span className="inline-block bg-yellow-200 text-black px-3 py-1 text-sm font-semibold mb-4">Simple Process</span>
                    <h2 className="text-4xl font-bold">
                        How To <span className="inline-block bg-green-300 px-1 -rotate-4">Scrape</span> Articles
                    </h2>
                    <p className="text-gray-700 mt-2">
                        Follow these three simple steps to scrape and save your articles easily
                    </p>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-3 gap-12 max-w-5xl mx-auto">
                    
                    <div className="text-center">
                        <div className="w-14 h-14 mx-auto rounded-full bg-blue-200 flex items-center justify-center text-xl font-bold mb-4 border-2 border-black">
                            1
                        </div>
                        <h3 className="font-semibold text-xl mb-2">Click Scrape</h3>
                        <p className="text-gray-600">Click the scrape button and enter the url you want to extract info from.</p>
                        <div className="bg-[#1a1a2e] text-blue-300 border-2 border-black transform rotate-1 mt-6 p-6">
                            Enter The Url
                        </div>
                    </div>

                    
                    <div className="text-center">
                        <div className="w-14 h-14 mx-auto rounded-full bg-green-200 flex items-center justify-center text-xl font-bold mb-4 border-2 border-black">
                            2
                        </div>
                        <h3 className="font-semibold text-xl mb-2">Wait for Results</h3>
                        <p className="text-gray-600">Our system processes the request and fetches the article for you.</p>
                        <div className="bg-[#1a1a2e] text-green-300 border-2 border-black transform -rotate-2 mt-6 p-6">
                           Wait..
                        </div>
                    </div>

                    
                    <div className="text-center">
                        <div className="w-14 h-14 mx-auto rounded-full bg-purple-200 flex items-center justify-center text-xl font-bold mb-4 border-2 border-black">
                            3
                        </div>
                        <h3 className="font-semibold text-xl mb-2">Press Done</h3>
                        <p className="text-gray-600">Once results are ready, press Done and you’re all set!</p>
                        <div className="bg-[#1a1a2e] text-purple-300 border-2 border-black transform rotate-1 mt-6 p-6">
                            Press Done
                        </div>
                    </div>
                </div>

                <div className="text-center mt-12">
                    <button className="bg-black text-white font-semibold px-6 py-3 rounded-md shadow-lg border-2 border-yellow-200 hover:scale-105 transition">
                        Try It Now
                    </button>
                </div>
            </section>
        </>
    );
}
