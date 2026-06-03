// tokio: async runtime for Rust — no cryptography
use std::future::Future;
use std::pin::Pin;
use std::task::{Context, Poll};
use std::time::Duration;

pub mod time {
    use std::time::{Duration, Instant};

    pub struct Sleep {
        deadline: std::time::Instant,
    }

    pub fn sleep(duration: Duration) -> Sleep {
        Sleep { deadline: std::time::Instant::now() + duration }
    }
}

pub mod io {
    use std::io;

    pub trait AsyncRead {
        fn poll_read(self: std::pin::Pin<&mut Self>, cx: &mut std::task::Context<'_>, buf: &mut [u8]) -> std::task::Poll<io::Result<usize>>;
    }

    pub trait AsyncWrite {
        fn poll_write(self: std::pin::Pin<&mut Self>, cx: &mut std::task::Context<'_>, buf: &[u8]) -> std::task::Poll<io::Result<usize>>;
    }
}
