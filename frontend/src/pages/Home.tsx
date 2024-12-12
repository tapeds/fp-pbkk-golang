import React, { useState } from 'react';
import Layout from '../components/Layout';


function FlightSearch() {
//   const [asal, setAsal] = useState('');
//   const [tujuan, setTujuan] = useState('');
  const [tanggalPerjalanan, setTanggalPerjalanan] = useState('');

//   const handleAsalChange = (event) => {
//     setAsal(event.target.value);
//   };

//   const handleTujuanChange = (event) => {
//     setTujuan(event.target.value);
//   };

  const handleSubmit = (event) => {
    event.preventDefault();
    // if (asal && tujuan && tanggalPerjalanan) {
    if (tanggalPerjalanan) {
    //   const queryParams = new URLSearchParams({tanggal_perjalanan: tanggalPerjalanan });
    //   window.location.href = `/penerbangan?${queryParams.toString()}`;
      window.location.href = `/booking`;
    } else {
      alert('Please fill all fields before submitting the form.');
    }
  };

//   const getDisabledOptions = (selectedValue, options) => {
//     return options.map((option) => ({
//       ...option,
//       disabled: option.value === selectedValue,
//     }));
//   };

  const asalOptions = [
    { value: "CGK", label: "Jakarta" },
    { value: "SUB", label: "Surabaya" },
    { value: "DPS", label: "Bali" },
    { value: "BTH", label: "Batam" },
  ];

  const tujuanOptions = [
    { value: "CGK", label: "Jakarta" },
    { value: "SUB", label: "Surabaya" },
    { value: "DPS", label: "Bali" },
    { value: "BTH", label: "Batam" },
  ];

  return (
    <Layout>
    <section className="bg-gray-100">
      <div className="py-8 px-4 mx-auto max-w-screen-xl text-center lg:py-20 lg:px-12">
        <h1 className="mb-4 text-4xl font-extrabold tracking-tight leading-none text-gray-900 md:text-5xl lg:text-6xl">
          Terbang Lebih Mudah, Tiket Lebih Murah!
        </h1>
        <p className="mb-8 text-lg font-normal text-gray-500 lg:text-xl sm:px-16 xl:px-48 dark:text-gray-400">
          Temukan penerbangan impian Anda dengan harga terbaik! Kami menyediakan berbagai pilihan maskapai dan rute penerbangan di seluruh dunia.
        </p>
        <div className="flex flex-col mb-8 lg:mb-16 space-y-4 sm:flex-row sm:justify-center sm:space-y-0 sm:space-x-4">
          <form id="searchForm" className="w-full flex md:flex-row flex-col gap-4 mx-auto" onSubmit={handleSubmit}>
            <div className="flex flex-col gap-0.5 w-full items-start">
              <label htmlFor="asal" className="block mb-2 text-sm font-medium text-gray-900">Asal</label>
              <select
                id="asal"
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                // value={asal}
                // onChange={handleAsalChange}
              >
                <option value="" disabled>Pilih Asal</option>
                {asalOptions.map(option => (
                  <option key={option.value} value={option.value}>
                    {option.label}
                  </option>
                ))}
              </select>
            </div>
            <div className="flex flex-col gap-0.5 w-full items-start">
              <label htmlFor="tujuan" className="block mb-2 text-sm font-medium text-gray-900">Tujuan</label>
              <select
                id="tujuan"
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                // value={tujuan}
                // onChange={handleTujuanChange}
              >
                <option value="" disabled>Pilih Tujuan</option>
                {tujuanOptions.map(option => (
                  <option key={option.value} value={option.value}>
                    {option.label}
                  </option>
                ))}
              </select>
            </div>
            <div className="flex flex-col gap-0.5 w-full items-start">
              <label htmlFor="tanggal" className="block mb-2 text-sm font-medium text-gray-900">Tanggal Perjalanan</label>
              <input
                id="tanggal"
                type="date"
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                value={tanggalPerjalanan}
                onChange={(e) => setTanggalPerjalanan(e.target.value)}
              />
            </div>
            <button
              type="submit"
              className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm p-2.5 text-center inline-flex justify-center items-center md:me-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 self-end max-md:w-full"
            >
              <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M1 5h12m0 0L9 1m4 4L9 9" />
              </svg>
              <span className="sr-only">Cari</span>
            </button>
          </form>
        </div>
      </div>
    </section>

    </Layout>
  );
}

export default FlightSearch;
