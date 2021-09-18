import Head from 'next/head'
import styles from './index.module.css'

export default function Home() {
  return (
    <div className={styles.container}>
      
      <h1 className="text-center">
        Selamat Datang ke Pegawai!
      </h1>
      <p className="mt-3 text-center">
        Ini hanya percobaan Kaenova untuk membuat suatu web apps dengan 
        Golang (Echo) [Backend] + MySQL [Database] + Docker - Docker Compose [Containerization] + Next.JS [Frontend]
      </p>
      <h2 className="mt-4 text-center pb-3">Penggunaan API</h2>
      <div className="ml-3 text-center grid">
        <a className={styles.list_container} href="#">Tampilkan Semua Pegawai</a>
        <a href="#" className={styles.list_container}> 
          Tampilkan Pegawai by ID 
          <p className="font-bold">Masukkan ID:</p>
          <input className="w-1/2 shadow rounded-sm" type="number" name="" id="" value="1"/>
        </a>
        <div>
          
        </div>
        <a className={styles.list_container} href="#">Tampilkan Semua Pegawai</a>
        <a className={styles.list_container} href="#">Tampilkan Semua Pegawai</a>
      </div>
    </div>
  )
}
