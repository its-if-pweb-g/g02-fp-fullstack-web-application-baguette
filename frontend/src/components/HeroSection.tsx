import React, { useState, useEffect } from 'react';

const images = [
  {
    src: '/heroroti.jpg',
    title: 'Rasa yang Menggugah, Roti yang Memikat',
    description: 'Eksplorasi produk roti terbaik kami, hanya untuk anda.',
  },
  {
    src: '/herochristmas.jpg',
    title: 'Promo Spesial Hari Natal',
    description: 'Diskon produk pilihan 25%, hanya pada 24-26 Desember 2024.',
  },
];

const HeroSection = () => {
  const [currentIndex, setCurrentIndex] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentIndex((prevIndex) => (prevIndex + 1) % images.length);
    }, 5000); 

    return () => clearInterval(interval);
  }, []);

  return (
    <section className="relative w-full h-screen overflow-hidden">
      {images.map((image, index) => (
        <div
          key={index}
          className={`absolute inset-0 w-full h-full transition-opacity duration-1000 ${
            index === currentIndex ? 'opacity-100' : 'opacity-0'
          }`}
        >
          <img
            src={image.src}
            alt={image.title}
            className="w-full h-full object-cover"
          />
          <div className="absolute inset-0 bg-black bg-opacity-40 flex flex-col items-center justify-center text-center px-4">
            <h1 className="text-4xl md:text-6xl text-white font-bold mb-4">
              {image.title}
            </h1>
            <p className="text-lg md:text-2xl text-white">
              {image.description}
            </p>
          </div>
        </div>
      ))}
    </section>
  );
};

export default HeroSection;
