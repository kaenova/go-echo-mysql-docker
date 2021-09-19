import Head from 'next/head'
import styles from './index.module.css'
import Link from 'next/link'

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
        <Link href="/pegawai">
        <a className={styles.list_container} href="#">Tampilkan Semua Pegawai</a>
        </Link>
        <a className={styles.list_container} href="#">Edit Pegawai</a>
        <a className={styles.list_container} href="#">Hapus Pegawai</a>
      </div>
    </div>
  )
}
